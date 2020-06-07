package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func EncryptToMd5(value *string) string {
	h := md5.New()
	h.Write([]byte(*value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncryptToSha1(value *string) string {
	h := sha1.New()
	h.Write([]byte(*value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncryptToSha256(value *string) string {
	h := sha256.New()
	h.Write([]byte(*value))
	return fmt.Sprintf("%x", h.Sum(nil))
}
