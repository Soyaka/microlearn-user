package utils

import (
	"math/rand"
	"time"
)

func GenerateOtp() string {
	const otpLength = 6
	const charset = "0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	otp := make([]byte, otpLength)
	for i := range otp {
		otp[i] = charset[r.Intn(len(charset))]
	}
	return string(otp)
}
