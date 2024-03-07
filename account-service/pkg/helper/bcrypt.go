package helper

import "golang.org/x/crypto/bcrypt"

func Hash(original string) (hashed string, err error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(original), 10)
	hashed = string(hashedByte)
	return
}
