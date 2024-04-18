package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func MD5Hash(password string) string {
	var hashPassword string

	h := md5.New()
	h.Write([]byte(password))
	hashPassword = fmt.Sprintf("%x", h.Sum(nil))

	return hashPassword
}

func SHA1Hash(password string) string {
	var hashPassword string

	h := sha1.New()
	h.Write([]byte(password))
	hashPassword = fmt.Sprintf("%x", h.Sum(nil))

	return hashPassword
}
