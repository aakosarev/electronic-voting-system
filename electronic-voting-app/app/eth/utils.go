package eth

import "github.com/ethereum/go-ethereum/accounts/keystore"

func GenerateNewAccount(keydir, password string) (address string, err error) {
	ks := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}
