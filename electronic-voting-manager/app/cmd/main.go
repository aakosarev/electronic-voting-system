package main

import (
	"context"
	"fmt"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/eth/voting"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	cfg := config.GetConfig()

	client, err := ethclient.Dial(cfg.Blockchain.URL)
	if err != nil {
		log.Fatal(err)
	}

	session, err := voting.NewSession(context.Background(), client, cfg)
	if err != nil {

	}

	//only for testing---------------------------------------------
	options := []string{"1", "2", "3"}
	title := "title"
	tString := "2023-03-26 15:04:05"
	t, _ := time.Parse("2006-01-02 15:04:05", tString)
	//--------------------------------------------------------------

	contractAddress, err := voting.RegisterVote(session, client, title, t, options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contractAddress)
}
