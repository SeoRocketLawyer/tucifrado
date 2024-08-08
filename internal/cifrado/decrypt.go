package cifrado

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func DecryptFile(password, filePath string) error {
	encryptedData, err := ioutil.ReadFile(filePath)
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

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	data, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	decryptedFilePath := filePath + ".dec"
	err = ioutil.WriteFile(decryptedFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
