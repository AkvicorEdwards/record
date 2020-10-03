package util

import (
	"encoding/hex"
	"log"
	"record/def"
	"record/hash"
)

func Encrypt(str string) string {
	if len(str) == 0 {
		return ""
	}
	data, err := hash.AesEncryptCBC([]byte(str), []byte(def.EncryptKey))
	if err != nil {
		log.Println(err)
		return ""
	}
	return hex.EncodeToString(data)
}

func Decrypt(str string) string {
	if len(str) == 0 {
		return ""
	}
	d, err := hex.DecodeString(str)
	if err != nil {
		log.Println(err)
		return ""
	}
	data, err := hash.AesDecryptCBC(d, []byte(def.EncryptKey))
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(data)
}