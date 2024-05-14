package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(bytes), err
}

func VerifyPassword(hashedPwd string, plainPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err 
}
