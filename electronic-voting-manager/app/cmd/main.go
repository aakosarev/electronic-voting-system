package main

func main() {
	/*cfg := config.GetConfig()

	client, err := ethclient.Dial(cfg.Blockchain.URL)
	if err != nil {
		log.Fatal(err)
	}

	session, err := voting.NewSession(context.Background(), client, cfg)
	if err != nil {
		log.Fatal(err)
	}

	//only for testing---------------------------------------------
	options := []string{"1", "2", "3"}
	title := "title"
	tString := "2023-03-26 15:04:05"
	t, _ := time.Parse("2006-01-02 15:04:05", tString)

	contractAddressStr, err := voting.CreateVoting(session, client, title, t, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contractAddressStr)

	contractAddress := common.HexToAddress(contractAddressStr)

	voterAddressStr := "0x23e4170970b57f335eD8362af9F97043770a4898"

	voterAddress := common.HexToAddress(voterAddressStr)

	balance, err := client.BalanceAt(context.Background(), voterAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	err = voting.GiveRightVote(session, client, cfg, contractAddress, voterAddress)
	if err != nil {
		log.Fatal(err)
	}

	balance, err = client.BalanceAt(context.Background(), voterAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
	//--------------------------------------------------------------*/

}
