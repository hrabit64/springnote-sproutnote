package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// ref : https://gist.github.com/donvito/efb2c643b724cf6ff453da84985281f8
func AesEncrypt(targetStr string, keyStr string) (encryptedStr string, err error) {

	key, _ := hex.DecodeString(keyStr)
	targetByte := []byte(targetStr)

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, targetByte, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

func AesDecrypt(encryptedStr string, keyStr string) (decryptedStr string, err error) {

	keyByte, _ := hex.DecodeString(keyStr)
	encryptedByte, _ := hex.DecodeString(encryptedStr)

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := encryptedByte[:nonceSize], encryptedByte[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
