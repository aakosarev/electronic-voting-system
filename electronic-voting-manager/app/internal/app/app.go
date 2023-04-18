package app

import (
	"context"
	"fmt"
	pb "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-manager/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/eth/voting"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/handler"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/service"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/storage"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/pkg/client/postgresql"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type App struct {
	cfg                  *config.Config
	grpcServer           *grpc.Server
	pgClient             *pgxpool.Pool
	votingManagerHandler pb.VotingManagerServer
	eclient              *ethclient.Client
	session              *voting.ContractSession
}

func NewApp(ctx context.Context, config *config.Config) (*App, error) {
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)

	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		return nil, err
	}

	eclient, err := ethclient.Dial(config.Blockchain.URL)
	if err != nil {
		return nil, err
	}

	session, err := voting.NewSession(context.Background(), eclient, config)
	if err != nil {
		return nil, err
	}

	votingStorage := storage.NewStorage(pgClient)
	votingService := service.NewService(votingStorage, session, eclient)

	votingManagerHandler := handler.NewHandler(
		votingService,
		pb.UnimplementedVotingManagerServer{},
	)

	return &App{
		cfg:                  config,
		pgClient:             pgClient,
		votingManagerHandler: votingManagerHandler,
		eclient:              eclient,
		session:              session,
	}, nil
}

func (a *App) startGRPC() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		return err
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)

	pb.RegisterVotingManagerServer(a.grpcServer, a.votingManagerHandler)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}

func (a *App) Run() error {
	return a.startGRPC()
}
