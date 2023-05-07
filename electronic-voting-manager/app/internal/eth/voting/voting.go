package voting

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/config"
	"github.com/aakosarev/electronic-voting-system/electronic-voting-manager/internal/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func NewSession(ctx context.Context, client *ethclient.Client, cfg *config.Config) (*ContractSession, error) {
	privateKey, err := crypto.HexToECDSA(cfg.Blockchain.PrivateKey)
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
	auth.Value = big.NewInt(cfg.Blockchain.WeiFounds)
	auth.GasLimit = uint64(cfg.Blockchain.GasLimit)
	auth.GasPrice = big.NewInt(cfg.Blockchain.GasPrice)

	return &ContractSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}, nil
}

func createContract(session *ContractSession, client *ethclient.Client, votingTitle string, votingEndTime time.Time) (string, error) {
	time.Sleep(2 * time.Second)
	nonce, err := client.PendingNonceAt(context.Background(), session.TransactOpts.From)
	if err != nil {
		return "", err
	}

	session.TransactOpts.Nonce = big.NewInt(int64(nonce))

	contractAddress, tx, instance, err := DeployContract(&session.TransactOpts, client, votingTitle, big.NewInt(votingEndTime.Unix()))
	if err != nil {
		return "", fmt.Errorf("failed to deploy the contract: %w", err)
	}

	timeout := 2 * time.Minute
	waitUntil := time.Now().Add(timeout)
	breakLoop := false
	for !breakLoop {
		receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
		if waitUntil.Sub(time.Now()) <= 0 {
			return "", fmt.Errorf("deploy transaction %s not mined, timing out after %v minutes", tx.Hash().Hex(), timeout)
		} else if receipt != nil {
			log.Printf("the contract deployment transaction with hash [ %s ] has been mined\n", tx.Hash().Hex())
			breakLoop = true
		}
	}

	session.Contract = instance

	return contractAddress.Hex(), nil
}

func loadContract(session *ContractSession, client *ethclient.Client, address common.Address) error {
	instance, err := NewContract(address, client)
	if err != nil {
		return fmt.Errorf("failed to load the contract: %w", err)
	}
	session.Contract = instance
	return nil
}

func CreateVoting(session *ContractSession, client *ethclient.Client, votingTitle string, votingEndTime time.Time, votingOptions []string) (string, error) {
	contractAddressString, err := createContract(session, client, votingTitle, votingEndTime)
	if err != nil {
		return "", fmt.Errorf("failed voting registration: %w", err)
	}

	contractAddress := common.HexToAddress(contractAddressString)

	err = addVotingOptions(session, client, contractAddress, votingOptions)
	if err != nil {
		return "", fmt.Errorf("failed addition of voting options: %w", err)
	}

	err = completeOptions(session, client, contractAddress)
	if err != nil {
		return "", fmt.Errorf("error complete options: %v", err)
	}

	return contractAddress.Hex(), nil
}

func addVotingOptions(session *ContractSession, client *ethclient.Client, address common.Address, votingOptions []string) error {
	err := loadContract(session, client, address)
	if err != nil {
		return err
	}

	var votingOptionsTransactionHashes []common.Hash
	for _, vo := range votingOptions {
		time.Sleep(2 * time.Second)
		nonce, err := client.PendingNonceAt(context.Background(), session.TransactOpts.From)
		if err != nil {
			log.Println(err)
			return err
		}

		session.TransactOpts.Nonce = big.NewInt(int64(nonce))

		tx, err := session.AddVotingOption(vo)
		if err != nil {
			log.Println(err)
			return err
		}

		votingOptionsTransactionHashes = append(votingOptionsTransactionHashes, tx.Hash())
	}

	timeout := 2 * time.Minute
	waitUntil := time.Now().Add(timeout)
	breakLoop := false
	for !breakLoop {
		allProcessed := true
		for _, trHash := range votingOptionsTransactionHashes {
			receipt, _ := client.TransactionReceipt(context.Background(), trHash)
			if receipt == nil {
				allProcessed = false
				break
			} else {
				log.Printf("the option addition transaction with hash [ %s ] has been mined\n", trHash.Hex())
			}
		}

		if waitUntil.Sub(time.Now()) <= 0 {
			return fmt.Errorf("add options transaction(s) not mined, timing out after %v minutes", timeout)
		} else if allProcessed {
			breakLoop = true
		}
	}

	return nil
}

func completeOptions(session *ContractSession, client *ethclient.Client, address common.Address) error {
	err := loadContract(session, client, address)
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)
	nonce, err := client.PendingNonceAt(context.Background(), session.TransactOpts.From)
	if err != nil {
		return err
	}

	session.TransactOpts.Nonce = big.NewInt(int64(nonce))

	tx, err := session.CompleteVotingOptions()
	if err != nil {
		return fmt.Errorf("failed complition voting options: %w", err)
	}

	timeout := 3 * time.Minute
	waitUntil := time.Now().Add(timeout)
	breakLoop := false
	for !breakLoop {
		receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
		if waitUntil.Sub(time.Now()) <= 0 {
			return fmt.Errorf("complete transaction %s not mined, timing out after %v minutes", tx.Hash().Hex(), timeout)
		} else if receipt != nil {
			log.Printf("the options completion transaction with hash [ %s ] has been mined\n", tx.Hash().Hex())
			breakLoop = true
		}
	}

	return nil
}

func RegisterAddressToVoting(session *ContractSession, client *ethclient.Client, cfg *config.Config, address common.Address, voterAddress common.Address) error {
	err := loadContract(session, client, address)
	if err != nil {
		return err
	}

	nonce, err := client.PendingNonceAt(context.Background(), session.TransactOpts.From)
	if err != nil {
		return err
	}

	session.TransactOpts.Nonce = big.NewInt(int64(nonce))

	tx, err := session.GiveRightToVote(voterAddress)
	if err != nil {
		return fmt.Errorf("failed giving right to vote: %w", err)
	}

	timeout := 2 * time.Minute
	waitUntil := time.Now().Add(timeout)
	breakLoop := false
	for !breakLoop {
		receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())
		if waitUntil.Sub(time.Now()) <= 0 {
			return fmt.Errorf("giving right to vote transaction %s not mined, timing out after %v minutes", tx.Hash().Hex(), timeout)
		} else if receipt != nil {
			log.Printf("the transaction of registering an address to voting with hash [ %s ] has been mined\n", tx.Hash().Hex())
			breakLoop = true
		}
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("error chainID: %w", err)
	}

	nonce, err = client.PendingNonceAt(context.Background(), session.TransactOpts.From)
	if err != nil {
		return fmt.Errorf("error pending nonce: %w", err)
	}

	value := big.NewInt(10000000000000000) // in wei

	tx = types.NewTransaction(nonce, voterAddress, value, session.TransactOpts.GasLimit, session.TransactOpts.GasPrice, nil)

	privateKey, err := crypto.HexToECDSA(cfg.Blockchain.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed conversion hex private key to ecdsa: %w", err)
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return fmt.Errorf("failed transaction signing: %w", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed send trans: %w", err)
	}

	waitUntil = time.Now().Add(timeout)
	breakLoop = false
	for !breakLoop {
		receipt, _ := client.TransactionReceipt(context.Background(), signedTx.Hash())
		if waitUntil.Sub(time.Now()) <= 0 {
			return fmt.Errorf("ether transfer transaction %s not mined, timing out after %v minutes", signedTx.Hash().Hex(), timeout)
		} else if receipt != nil {
			log.Printf("the ether transfer transaction with hash [ %s ] has been mined\n", signedTx.Hash().Hex())
			breakLoop = true
		}
	}

	return nil
}

func GetInformationAboutVoting(session *ContractSession, client *ethclient.Client, address common.Address) (*model.InformationAboutVoting, error) {
	err := loadContract(session, client, address)
	if err != nil {
		return nil, err
	}

	votingTitle, _ := session.GetVotingTitle()
	numberRegisteredVotersBigInt, _ := session.GetNumberRegisteredVoters()
	optionsCompleted, _ := session.GetOptionsCompleted()
	endTimeBigInt, _ := session.GetVotingEndTime()
	endTime := time.Unix(endTimeBigInt.Int64(), 0)
	numberOptionsBigInt, _ := session.GetVotingOptionsLength()

	options := make(map[int64]*model.Option, numberOptionsBigInt.Int64())

	for i := int64(0); i < numberOptionsBigInt.Int64(); i++ {
		name, _ := session.GetNameVotingOption(big.NewInt(i))
		numberVotes, _ := session.GetNumberVotesVotingOption(big.NewInt(i))
		options[i] = &model.Option{
			Name:        name,
			NumberVotes: numberVotes.Int64(),
		}
	}

	return &model.InformationAboutVoting{
		Title:                  votingTitle,
		OptionsCompleted:       optionsCompleted,
		NumberRegisteredVoters: numberRegisteredVotersBigInt.Int64(),
		EndTime:                endTime,
		Options:                options,
	}, nil
}
