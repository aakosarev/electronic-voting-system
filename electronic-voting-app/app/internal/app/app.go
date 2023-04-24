package app

import (
	"context"
	"fmt"
	pb "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-app/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/config"
	mygrpc "github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/handler/grpc"
	myhttp "github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/handler/http"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/internal/storage"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-app/pkg/client/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg                  *config.Config
	grpcServer           *grpc.Server
	router               *httprouter.Router
	httpServer           *http.Server
	pgClient             *pgxpool.Pool
	votingAppGRPCHandler pb.VotingAppServer
}

func NewApp(ctx context.Context, config *config.Config) (*App, error) {
	router := httprouter.New()

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)

	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		return nil, err
	}

	userStorage := storage.NewUserStorage(pgClient)

	votingAppHTTPHandler := myhttp.NewHandler(userStorage)
	votingAppHTTPHandler.Register(router)

	votingAppGRPCHandler := mygrpc.NewHandler(
		userStorage,
		pb.UnimplementedVotingAppServer{},
	)

	return &App{
		cfg:                  config,
		router:               router,
		pgClient:             pgClient,
		votingAppGRPCHandler: votingAppGRPCHandler,
	}, nil
}

func (a *App) startGRPC() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		return err
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)

	pb.RegisterVotingAppServer(a.grpcServer, a.votingAppGRPCHandler)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}

func (a *App) startHTTP() error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Handler:      a.router,
		WriteTimeout: 30 * time.Second, //TODO get in cfg
		ReadTimeout:  30 * time.Second,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP()
	})
	grp.Go(func() error {
		return a.startGRPC()
	})
	return grp.Wait()

	/*err := a.startHTTP()
	if err != nil {
		return err
	}

	err = a.startGRPC()
	if err != nil {
		return err
	}
	return nil*/
}
