package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	bao "github.com/openbao/openbao/api/v2"
)

const (
	vaultAddr   = "http://127.0.0.1:8200"
	vaultPath   = "secret/data/my-aes-key" // OpenBao KV v2 path
	s3Bucket    = "tinkr18k-encryption-client"
	s3ObjectKey = "encrypted-demo.txt"
)

func main() {
	ctx := context.TODO()

	// Get AES key from bao
	key, err := getAESKeyFromVault()
	if err != nil {
		log.Fatalf("failed to get AES key from bao: %v", err)
	}
	fmt.Println("🔑 AES key retrieved from bao")

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	s3Client := s3.NewFromConfig(cfg)

	// Encrypt and upload
	plaintext := []byte("This is a secret message to store encrypted in S3.")
	encrypted, nonce, err := encryptAESGCM(plaintext, key)
	if err != nil {
		log.Fatalf("encryption failed: %v", err)
	}
	payload := append(nonce, encrypted...)

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3ObjectKey),
		Body:   bytes.NewReader(payload),
	})
	if err != nil {
		log.Fatalf("S3 upload failed: %v", err)
	}
	fmt.Println("✅ Encrypted file uploaded to S3")

	// Download and decrypt
	resp, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3ObjectKey),
	})
	if err != nil {
		log.Fatalf("failed to download from S3: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read failed: %v", err)
	}

	nonceSize := 12
	nonceFromS3 := data[:nonceSize]
	cipherFromS3 := data[nonceSize:]

	decrypted, err := decryptAESGCM(cipherFromS3, nonceFromS3, key)
	if err != nil {
		log.Fatalf("decryption failed: %v", err)
	}
	fmt.Println("🔓 Decrypted content:", string(decrypted))
}

// getAESKeyFromVault retrieves and decodes AES key from OpenBao
func getAESKeyFromVault() ([]byte, error) {
	client, err := bao.NewClient(&bao.Config{Address: vaultAddr})
	if err != nil {
		return nil, err
	}
	client.SetToken("localenv")

	// Authenticate via env VAULT_TOKEN or ~/.bao-token
	secret, err := client.Logical().Read(vaultPath)
	if err != nil {
		return nil, err
	}
	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("no secret found at %s", vaultPath)
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected secret format")
	}
	rawKey, ok := data["key"].(string)
	if !ok {
		return nil, fmt.Errorf("key not found in bao")
	}
	return base64.StdEncoding.DecodeString(rawKey)
}

func encryptAESGCM(plaintext, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, nil, err
	}
	return aesgcm.Seal(nil, nonce, plaintext, nil), nonce, nil
}

func decryptAESGCM(ciphertext, nonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
