package main

import (
	"context"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/app"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/config"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
