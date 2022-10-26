package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	md5 := md5.Sum([]byte("5d4b7a49339eb3191688116224871548Mysoft!@#$%^&"))

	sha1 := sha1.Sum([]byte("5d4b7a49339eb3191688116224871548Mysoft!@#$%^&"))

	key := GetByteArray(converArr1(sha1), 32)

	iv := converArr2(md5)

	fmt.Println(NewAESCrypt(key, iv, AESModeCBC).Encrypt("{\"user_code\":\"huxz011111111111111111\"}"))

}

const (
	AESModeCBC AESMode = "CBC"
)

type AESMode string

type AESCrypt struct {
	Key  []byte
	IV   []byte
	Mode AESMode
}

func NewAESCrypt(key []byte, iv []byte, mode AESMode) AESCrypt {
	return AESCrypt{Key: key, IV: iv, Mode: mode}
}

func (a AESCrypt) Encrypt(value string) (result string, err error) {
	originData := []byte(value)
	var encrypted []byte
	switch a.Mode {
	case AESModeCBC:
		encrypted, err = a.aesEncryptCBC(originData)
	}
	if err != nil {
		return
	}

	result = base64.StdEncoding.EncodeToString(encrypted)

	fmt.Println(result)

	return
}

func (a AESCrypt) aesEncryptCBC(originData []byte) (encrypted []byte, err error) {
	var myCipher cipher.Block
	myCipher, err = aes.NewCipher(a.Key)
	if err != nil {
		return
	}
	blockSize := myCipher.BlockSize()                   // 获取秘钥块的长度
	originData = pkcs5Padding(originData, blockSize)    // 补全码
	blockMode := cipher.NewCBCEncrypter(myCipher, a.IV) // 加密模式
	encrypted = make([]byte, len(originData))           // 创建数组
	blockMode.CryptBlocks(encrypted, originData)
	return
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func converArr1(array [20]byte) []byte {
	s := []byte{}
	for _, elem := range array {
		s = append(s, elem)
	}
	return s
}

func converArr2(array [16]byte) []byte {
	s := []byte{}
	for _, elem := range array {
		s = append(s, elem)
	}
	return s
}

func GetByteArray(src []byte, destLen int) []byte {
	dest := make([]byte, destLen)
	p := 0
	for p < destLen {
		for _, b := range src {
			if p >= destLen {
				return dest
			} else {
				dest[p] = b
				p++
			}
		}
	}
	return dest
}
