package tools

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
)

var salt = []byte{0x7a, 0x73, 0x64, 0x6c, 0x7a, 0x6a, 0x71, 0x61}

func EncryptPwd(source string) string {
	dk, _ := scrypt.Key([]byte(source), salt, 1<<15, 8, 1, 32)
	return base64.StdEncoding.EncodeToString(dk)
}
