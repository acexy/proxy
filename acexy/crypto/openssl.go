package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/pbkdf2"
)

func DecryptOpenSSL(cipherData []byte, password string) ([]byte, error) {
	if len(cipherData) < 16 {
		return nil, errors.New("data too short")
	}
	header := cipherData[:16]
	if !bytes.HasPrefix(header, []byte("Salted__")) {
		return nil, errors.New("invalid openssl salted format")
	}
	salt := header[8:]

	// 派生 Key 和 IV
	key := pbkdf2.Key([]byte(password), salt, 10000, 32+16, sha256.New)
	aesKey := key[:32]
	iv := key[32:]

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	ciphertext := cipherData[16:]
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除 PKCS#7 padding
	padding := int(plaintext[len(plaintext)-1])
	if padding == 0 || padding > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}
	return plaintext[:len(plaintext)-padding], nil
}
