package axa_security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func SHA256EncryptPassword(password string) string{
	hash := sha256.New()
	hash.Write([]byte(password))
	pass_sum := hash.Sum(nil)
	return hex.EncodeToString(pass_sum)
}

// THE CODE BELOW WAS MORE OR LESS COPIED
// I OWE MOST OF THE CREDIT FOR THIS CODE TO:
// Aziza Kasenova
// AND TO HER ARTICLE:
// https://medium.com/insiderengineering/aes-encryption-and-decryption-in-golang-php-and-both-with-full-codes-ceb598a34f41
func EncryptData(data string) (string, error){
	key := GetEnvironmentVarValue("AXADB_AES_KEY")
	iv := GetEnvironmentVarValue("AXADB_AES_IV")

	var plainTextBlock []byte
	dataLength := len(data)

	if dataLength % 16 != 0 {
		extendBlock := 16 - (dataLength % 16)
		plainTextBlock = make([]byte, dataLength+extendBlock)
		copy(plainTextBlock[dataLength:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, dataLength)
	}

	copy(plainTextBlock, data)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil

}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func DecryptData(encrypted string) ([]byte, error){
	key := GetEnvironmentVarValue("AXADB_AES_KEY")
	iv := GetEnvironmentVarValue("AXADB_AES_IV")

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}