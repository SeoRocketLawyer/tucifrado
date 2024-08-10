package cifrado

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func DecryptFile(password, filePath string) error {
	key := deriveKey(password)
	encryptedData, err := os.ReadFile(filePath)
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

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	data, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Obtener el nombre original del archivo (sin la extensión .enc)
	originalFilePath := filePath[:len(filePath)-len(filepath.Ext(filePath))]

	// Reemplazar el archivo encriptado con el archivo desencriptado
	err = os.WriteFile(originalFilePath, data, 0644)
	if err != nil {
		return err
	}

	// Opcional: Eliminar el archivo encriptado después de desencriptar
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
