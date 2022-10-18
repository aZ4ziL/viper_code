/*
This is implement the user model
*/

package models

import (
	"time"

	"github.com/aZ4ziL/viper_code/utils"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	FirstName  string    `gorm:"size:100"`
	LastName   string    `gorm:"size:100"`
	Username   string    `gorm:"size:100;unique;index"`
	Email      string    `gorm:"size:100;unique;index"`
	Password   string    `gorm:"size:255"`
	IsAdmin    bool      `gorm:"default:0"`
	IsActive   bool      `gorm:"default:1"`
	LastLogin  time.Time `gorm:"null"`
	DateJoined time.Time `gorm:"autoCreateTime"`
}

// CreateNewUser create new user if user is exists return error
func CreateNewUser(user *User) error {
	user.Password = utils.EncryptionPassword(user.Password)
	err = db.Create(user).Error
	return err
}

// GetUserByUsername get user by passing the username.
// If user by the username is not found return error not found.
func GetUserByUsername(username string) (User, error) {
	var user User
	err = db.Model(&User{}).Where("username = ?", username).First(user).Error
	return user, err
}

// GetUserByID get user by passing the ID.
// If user by the ID is not found return error not found.
func GetUserByID(id uint) (User, error) {
	var user User
	err = db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return user, err
}
