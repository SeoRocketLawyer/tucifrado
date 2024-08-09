package cifrado

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
)

func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func EncryptFile(password, filePath string) error {
	key := deriveKey(password)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
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
	err = os.WriteFile(encryptedFilePath, encryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
