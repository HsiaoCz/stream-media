package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var my_secert = "xiaofanyi"

// use md5 to Encrypt Password
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(my_secert))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
