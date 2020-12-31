package handle

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES encryption function.
func encrypt(orig []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	orig = padding(orig, blockSize)
	cryted := make([]byte, len(orig))
	blockMode.CryptBlocks(cryted, orig)
	return []byte(base64.StdEncoding.EncodeToString(cryted))
}

// AES decryption function.
func decrypt(cryted []byte, key []byte) []byte {
	cryted, _ = base64.StdEncoding.DecodeString(string(cryted))
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	orig := make([]byte, len(cryted))
	blockMode.CryptBlocks(orig, cryted)
	orig = unpadding(orig)
	return orig
}

// The length of the key is 16, but to ensure that each key corresponds to the
// data one-to-one, some characters must be added or some characters deleted.

func padding(text []byte, size int) []byte {
	length := size - len(text)%size
	postfix := bytes.Repeat([]byte{byte(length)}, length)
	return append(text, postfix...)
}

func unpadding(text []byte) []byte {
	length := len(text)
	postfix := int(text[length-1])
	return text[:(length - postfix)]
}
