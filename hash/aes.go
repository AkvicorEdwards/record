package hash

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte, err error) {
	if !(len(key)==16 || len(key)==24 || len(key)==32) {
		return []byte(""), errors.New("aes key must be either 16, 24, or 32 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""),err
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted, nil
}

func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	if !(len(key)==16 || len(key)==24 || len(key)==32) {
		return []byte(""), errors.New("aes key must be either 16, 24, or 32 bytes")
	}
	block, err := aes.NewCipher(key)                              // 分组秘钥
	if err != nil {
		return []byte(""),err
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted, nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
