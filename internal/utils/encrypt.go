package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts cipher text string into plain text string using AES algorithm.
func Decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func AesEncoding(src string, key []byte) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher(key)
	if err != nil {
		return src
	}
	// 密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) // 由于字节长度不够，所以要进行字节的填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64 编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}
