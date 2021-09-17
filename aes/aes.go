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
		log.Fatalf("USAGE: %s key text1 text2 \n", os.Args[0])
		return
	}
	key := os.Args[1]
	if strings.Contains(os.Args[0], "aes-decrypt") {
		for _, text := range os.Args[2:] {
			source, err := decrypt(text, key)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(source)
		}
	} else {
		for _, text := range os.Args[2:] {
			txt, err := encrypt(text, key)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(txt)
		}
	}
}

// Encrypt with AES/CBC/PKCS5Padding and base64-encoded
func encrypt(plainText string, key string) (secretText string, err error) {
	iv := []byte(_IV)
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCEncrypter(block, iv)
	var source = pkcs5pad([]byte(plainText), 16)
	var dst = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	secretText = base64.StdEncoding.EncodeToString(dst)
	return
}

func decrypt(secretText string, key string) (plainText string, err error) {
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
	dst := make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	plainText = string(pkcs5unpad(dst))
	return
}

func pkcs5pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5unpad(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
