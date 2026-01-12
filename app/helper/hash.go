package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckPassword(HashPassword string, password string) error {
	errCheck := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))

	return errCheck
}
