package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func HashWithSha256(plaintext string) (string, error) {
	h := sha256.New()
	if _, err := io.WriteString(h, plaintext); err != nil {
		return "", err
	}
	r := h.Sum(nil)
	return hex.EncodeToString(r), nil
}

func NewCipherBlock(key string) (cipher.Block, error) {
	hashedKey, err := HashWithSha256(key)
	if err != nil {
		return nil, err
	}
	bs, err := hex.DecodeString(hashedKey)
	if err != nil {
		return nil, err
	}
	return aes.NewCipher(bs[:])
}

func Pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func Pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}

	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		fmt.Println("here", n)
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			fmt.Println("hereeee")
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return b[:len(b)-n], nil
}

// Encrypt encrypts a plaintext
func Encrypt(key, plaintext string) (string, error) {
	block, err := NewCipherBlock(key)
	if err != nil {
		return "", err
	}

	//pad plaintext
	ptbs, _ := Pkcs7Pad([]byte(plaintext), block.BlockSize())

	if len(ptbs)%aes.BlockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}

	ciphertext := make([]byte, len(ptbs))

	//create an Initialisation vector which is the length of the block size for AES
	var iv []byte = make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	//Encrypt plaintext
	mode.CryptBlocks(ciphertext, ptbs)

	//concatenate initialisation vector and ciphertext
	return hex.EncodeToString(iv) + ":" + hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext
func Decrypt(key, ciphertext string) (string, error) {
	block, err := NewCipherBlock(key)
	if err != nil {
		return "", err
	}

	//split ciphertext into initialisation vector and actual ciphertext
	ciphertextParts := strings.Split(ciphertext, ":")
	iv, err := hex.DecodeString(ciphertextParts[0])
	if err != nil {
		return "", err
	}
	ciphertextbs, err := hex.DecodeString(ciphertextParts[1])
	if err != nil {
		return "", err
	}

	if len(ciphertextParts[1]) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(ciphertextParts[1])%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt cipher text
	mode.CryptBlocks(ciphertextbs, ciphertextbs)

	// Unpad ciphertext
	ciphertextbs, err = Pkcs7Unpad(ciphertextbs, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(ciphertextbs), nil
}
