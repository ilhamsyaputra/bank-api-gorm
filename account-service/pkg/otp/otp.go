package otp

import (
	"math/rand"
	"time"
)

func GenerateOtp() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(9999-1000+1) + 1000
}
