package keystorage

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/cryptoballot/rsablind"
	"os"
	"strconv"
)

type KeyStorage struct{}

func NewKeyStorage() *KeyStorage {
	return &KeyStorage{}
}

func (ks *KeyStorage) GenerateKeyPairForVotingID(votingID int32) error {
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

func (ks *KeyStorage) GetPublicKeyBytesForVotingID(votingID int32) ([]byte, error) {
	wd, _ := os.Getwd()
	filenamePublicKey := fmt.Sprintf("%s/internal/keystorage/keys/voting_%s_public.pem", wd, strconv.Itoa(int(votingID)))

	publicKeyBytes, err := os.ReadFile(filenamePublicKey)
	if err != nil {
		return nil, err
	}

	return publicKeyBytes, nil
}

func (ks *KeyStorage) getPrivateKeyForVotingID(votingID int32) (*rsa.PrivateKey, error) {
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

func (ks *KeyStorage) SignBlindedToken(blindedToken []byte, votingID int32) ([]byte, error) {
	privateKey, err := ks.getPrivateKeyForVotingID(votingID)
	if err != nil {
		return nil, err
	}

	signedBlindedToken, err := rsablind.BlindSign(privateKey, blindedToken)
	if err != nil {
		return nil, err
	}

	return signedBlindedToken, nil
}

func (ks *KeyStorage) VerifySignature(signedToken, token []byte, votingID int32) (bool, error) {

	publicKeyBytes, err := ks.GetPublicKeyBytesForVotingID(votingID)
	if err != nil {
		return false, err
	}

	publicKeyPEM, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyPEM.Bytes)
	if err != nil {
		return false, err
	}

	if err = rsablind.VerifyBlindSignature(publicKey, token, signedToken); err != nil {
		return false, err
	}

	return true, nil
}
