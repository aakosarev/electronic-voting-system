package main

import (
	"fmt"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-verifier/internal/keystorage"
	"log"
)

func main() {
	err := keystorage.GenerateKeyPairForVotingID(12345)
	if err != nil {
		log.Fatal(err)
	}

	s, err := keystorage.GetPublicKey(12345)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
