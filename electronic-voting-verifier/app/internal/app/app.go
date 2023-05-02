package app

import (
	"context"
	"fmt"
	pbvv "github.com/aakosarev/electronic-voting-system/contracts/gen/go/electronic-voting-verifier/v1"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/handler"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/keystorage"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/storage"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/pkg/client/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type App struct {
	cfg                   *config.Config
	grpcServer            *grpc.Server
	pgClient              *pgxpool.Pool
	votingVerifierHandler pbvv.VotingVerifierServer
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

	votingManagerConn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.VotingManagerGRPC.IP, config.VotingManagerGRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	storage := storage.NewStorage(pgClient)
	kstorage := keystorage.NewKeyStorage()

	votingVerifierHandler := handler.NewHandler(
		kstorage,
		storage,
		pbvv.UnimplementedVotingVerifierServer{},
		votingManagerConn,
	)

	return &App{
		cfg:                   config,
		pgClient:              pgClient,
		votingVerifierHandler: votingVerifierHandler,
	}, nil
}

func (a *App) startGRPC() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		return err
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)

	pbvv.RegisterVotingVerifierServer(a.grpcServer, a.votingVerifierHandler)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}

func (a *App) Run() error {
	return a.startGRPC()
}
