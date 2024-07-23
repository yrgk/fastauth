package storage

import (
	"fastauth/auth"
	"fastauth/structs"

	"gorm.io/gorm"
)

func CreateUser(userPayload structs.RegisterRequest, db *gorm.DB) error {
	newUser := User{
		Username: userPayload.Username,
		Email: userPayload.Email,
		Password: userPayload.Password,
	}

	if err := db.Create(&newUser); err != nil {
		return err.Error
	}
	
	return nil
}

// Getting user from db
func GetUser(email string, db *gorm.DB) User {
	var user User
	db.Raw("SELECT * FROM users WHERE users.email = ?", email).Scan(&user)

	return user
}

// Verifying user
func VerifyUser(userPayload structs.LoginRequest, db *gorm.DB) bool {
	user := GetUser(userPayload.Email, db)
	if user.ID == 0 {
		return false
	}

	return auth.VerifyPassword(userPayload.Password, user.Password)
}