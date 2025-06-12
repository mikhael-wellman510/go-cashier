package utils

import "golang.org/x/crypto/bcrypt"

// Todo -> hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// todo -> Cek password hashing
func CheckPasswordHash(hashPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))

	return err == nil
}
