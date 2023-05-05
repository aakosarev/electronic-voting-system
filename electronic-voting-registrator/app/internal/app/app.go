package app

import (
	"context"
	"fmt"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-registrator/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-registrator/internal/handler"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg        *config.Config
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	router := httprouter.New()

	votingManagerConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.VotingManagerGRPC.IP, cfg.VotingManagerGRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	votingAppConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.VotingAppGRPC.IP, cfg.VotingAppGRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	registratorHandler := handler.NewHandler(votingManagerConn, votingAppConn)

	registratorHandler.Register(router)

	return &App{
		cfg:    cfg,
		router: router,
	}, nil
}

func (a *App) startHTTP() error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Handler:      a.router,
		WriteTimeout: 5 * time.Minute, //TODO get in cfg
		ReadTimeout:  5 * time.Minute,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (a *App) Run() error {
	return a.startHTTP()
}
