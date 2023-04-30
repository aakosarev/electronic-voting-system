package keystorage

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strconv"
)

func GenerateKeyPairForVotingID(votingID int32) error {
	wd, _ := os.Getwd()
	filenamePrivateKey := fmt.Sprintf("%s/internal/keystorage/keys/voting_%s_private.pem", wd, strconv.Itoa(int(votingID)))
	filenamePublicKey := fmt.Sprintf("%s/internal/keystorage/keys/voting_%s_public.pem", wd, strconv.Itoa(int(votingID)))

	_, err1 := os.Stat(filenamePrivateKey)
	_, err2 := os.Stat(filenamePublicKey)

	if (os.IsNotExist(err1) && err2 == nil) || (os.IsNotExist(err2) && err1 == nil) || (os.IsNotExist(err1) && os.IsNotExist(err2)) {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return err
		}

		privateKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}

		privateFile, err := os.Create(filenamePrivateKey)
		if err != nil {
			return err
		}
		defer privateFile.Close()

		err = pem.Encode(privateFile, privateKeyPEM)
		if err != nil {
			return err
		}

		publicKeyPEM := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		}

		publicFile, err := os.Create(filenamePublicKey)
		if err != nil {
			return err
		}
		defer publicFile.Close()

		err = pem.Encode(publicFile, publicKeyPEM)
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	return nil
}

func GetPublicKeyForVotingID(votingID int32) (*rsa.PublicKey, error) {
	wd, _ := os.Getwd()
	filenamePublicKey := fmt.Sprintf("%s/internal/keystorage/keys/voting_%s_public.pem", wd, strconv.Itoa(int(votingID)))

	publicKeyBytes, err := os.ReadFile(filenamePublicKey)
	if err != nil {
		return nil, err
	}

	publicKeyPEM, _ := pem.Decode(publicKeyBytes)

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyPEM.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func GetPrivateKeyForVotingID(votingID int32) (*rsa.PrivateKey, error) {
	wd, _ := os.Getwd()
	filenamePrivateKey := fmt.Sprintf("%s/internal/keystorage/keys/voting_%s_private.pem", wd, strconv.Itoa(int(votingID)))

	privateKeyBytes, err := os.ReadFile(filenamePrivateKey)
	if err != nil {
		return nil, err
	}

	privateKeyPEM, _ := pem.Decode(privateKeyBytes)

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPEM.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
