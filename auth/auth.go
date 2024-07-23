package auth

import (
	"fastauth/structs"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Creating access token
func CreateToken(payload structs.LoginRequest, secretKey string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			// "username": payload.Username,
			"email": payload.Email,
			"password": payload.Password,
		})

	token, err := t.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return token
}

// Hashing plain password using
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Verifying password
func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}