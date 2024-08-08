package cifrado

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func EncryptFile(password, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	encryptedData := gcm.Seal(nonce, nonce, data, nil)
	encryptedFilePath := filePath + ".enc"
	err = ioutil.WriteFile(encryptedFilePath, encryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
