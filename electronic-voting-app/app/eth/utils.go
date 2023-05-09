package eth

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

/*func GenerateNewAccount(keydir, password string) (address string, err error) {
	ks := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)

	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}*/

func GenerateNewKeyPair() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, privateKeyString, nil
}

/*func GetAddressesByPrivateKeys(privateKeys []string) ([]string, error) {
	var addresses []string

	for _, privateKey := range privateKeys {
		privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, err
		}

		publicKey := privateKeyECDSA.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("error casting public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		addresses = append(addresses, address)
	}

	return addresses, nil
}*/

func GetAddressesByPrivateKeys(privateKeys map[int32]string) (map[int32]string, error) {
	addresses := make(map[int32]string)

	for key, privateKey := range privateKeys {
		privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, err
		}

		publicKey := privateKeyECDSA.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("error casting public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		addresses[key] = address
	}

	return addresses, nil
}
