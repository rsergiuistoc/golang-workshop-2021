package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model, defines the user state
type User struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`

	FirstName	string		`gorm:"size:150" json:"first_name"`
	LastName	string		`gorm:"size:150" json:"last_name"`
	Email		string		`gorm:"size:254; not null;unique" json:"email" binding:"required"`
	Password	string		`gorm:"size:128" json:"password" binding:"required"`

	Todo		[]Todo		`gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE" json:"todos"`
}

// Hash generates a hashed password from an plain text one
func Hash(password string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword verifies if the incoming password
// matches with the hashed one.
func CheckPassword(hashedPassword ,password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User)BeforeCreate(tx *gorm.DB) (err error){
	hashedPassword, _ := Hash(u.Password)
	u.Password = string(hashedPassword)

	return
}