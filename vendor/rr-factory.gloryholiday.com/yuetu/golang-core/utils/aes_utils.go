package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

/**
PKCS5包装
*/
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
解包装
*/
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func Encrypt(key, text []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher(key)
	content := PKCS5Padding(text, aesBlockEncrypt.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]

	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, iv)
	aesEncrypt.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	decrypted := make([]byte, len(text))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]

	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, iv)
	aesDecrypt.CryptBlocks(decrypted, text)
	return PKCS5Trimming(decrypted), nil
}

func EncryptBase64(key, text []byte) (string, error) {
	encrypted, err := Encrypt(key, text)

	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(encrypted)

	return base64Str, nil
}

func DecryptBase64(key []byte, encodeString string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		return []byte(""), err
	}

	result, err := Decrypt(key, decodeBytes)
	if err != nil {
		return nil, err
	}

	return result, nil
}
