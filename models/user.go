package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

// BeforeSave hash password before saving
func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password is correct
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// CreateUser creates a new user in the database
func CreateUser(user *User) error {
	err := user.BeforeSave()
	if err != nil {
		return err
	}
	return db.Create(user).Error
}

// AuthenticateUser authenticates a user during login
func AuthenticateUser(user *User) error {
	var dbUser User
	err := db.Where("username = ?", user.Username).First(&dbUser).Error
	if err != nil {
		return err
	}

	err = dbUser.CheckPassword(user.Password)
	return err
}

// GetUserByID retrieves a user by ID
func GetUserByID(id uint) (*User, error) {
	user := &User{}
	err := db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
