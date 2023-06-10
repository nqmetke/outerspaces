package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not nulll;" json:"password"`
	Bio      string
	Posts    []Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	Ups      int
	Downs    int
	Space    Space
	UserID   uint
	SpaceID  uint
	ParentID uint
	UppedBy  []User `gorm:"many2many:post_uppedby;"`
	DownedBy []User `gorm:"many2many:post_downedBy;"`
	Comments []Post `gorm:"foreignkey:ParentID"`
}

type Space struct {
	gorm.Model
	Title       string
	Description string
	Users       []User `gorm:"many2many:space_users;"`
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
	u.ID = 0
}
