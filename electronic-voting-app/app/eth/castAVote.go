package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func NewSession(ctx context.Context, client *ethclient.Client, privateKeyHex string) (*ContractSession, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed conversion hex private key to ecdsa: %w", err)
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed a getting public key from a private key: %w", err)
	}

	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("failed pending nonce: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed retrieves chainID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed creation transaction signer: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(10000000)
	auth.GasPrice = big.NewInt(4000000000)

	return &ContractSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}, nil
}

func loadContract(session *ContractSession, client *ethclient.Client, address common.Address) error {
	instance, err := NewContract(address, client)
	if err != nil {
		return fmt.Errorf("failed to load the contract: %w", err)
	}
	session.Contract = instance
	return nil
}

func CastAVote(session *ContractSession, client *ethclient.Client, address common.Address, idx int32) (string, error) {
	err := loadContract(session, client, address)
	if err != nil {
		return "", err
	}

	time.Sleep(2 * time.Second)
	nonce, err := client.PendingNonceAt(context.Background(), session.TransactOpts.From)
	if err != nil {
		return "", err
	}

	session.TransactOpts.Nonce = big.NewInt(int64(nonce))

	tx, err := session.Vote(big.NewInt(int64(idx)))
	if err != nil {
		return "", fmt.Errorf("failed cast a vote: %w", err)
	}

	timeout := 2 * time.Minute
	waitUntil := time.Now().Add(timeout)
	breakLoop := false
	for !breakLoop {
		receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
		if waitUntil.Sub(time.Now()) <= 0 {
			return "", fmt.Errorf("the cast a vote transaction %s not mined, timing out after %v minutes", tx.Hash().Hex(), timeout)
		} else if receipt != nil {
			log.Printf("the cast a vote transaction with hash [ %s ] has been mined\n", tx.Hash().Hex())
			breakLoop = true
		}
	}

	return tx.Hash().Hex(), nil
}
