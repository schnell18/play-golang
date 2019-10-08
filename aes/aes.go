package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

const _IV = "aabbccddeeffgghh"

func main() {
	if len(os.Args) < 3 {
		fmt.Errorf("USAGE: %s key text1 text2 \n", os.Args[0])
		return
	}
	key := os.Args[1]
	if strings.Contains(os.Args[0], "aes-decrypt") {
		for _, text := range os.Args[2:] {
			source, err := AESBase64Decrypt(text, key)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(source)
		}
	} else {
		for _, text := range os.Args[2:] {
			txt, err := AESBase64Encrypt(text, key)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(txt)
		}
	}
}

// Encrypt with AES/CBC/PKCS5Padding and base64-encoded
func AESBase64Encrypt(plainText string, key string) (secretText string, err error) {
	iv := []byte(_IV)
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCEncrypter(block, iv)
	var source []byte = _PKCS5Padding([]byte(plainText), 16)
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	secretText = base64.StdEncoding.EncodeToString(dst)
	return
}

func AESBase64Decrypt(secretText string, key string) (plainText string, err error) {
	iv := []byte(_IV)
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCDecrypter(block, iv)

	var source []byte
	if source, err = base64.StdEncoding.DecodeString(secretText); err != nil {
		log.Println(err)
		return
	}
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	plainText = string(_PKCS5Unpadding(dst))
	return
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
