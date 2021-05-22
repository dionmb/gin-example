package models

import "gin_example/libs/auth"

type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null"`
	PasswordDigest string
	Activated bool `gorm:"default:true"`
	Jti string `gorm:"uniqueIndex;not null;default:uuid_generate_v4()"`
}

func (u *User) SetPassword(password string) {
	u.PasswordDigest = auth.EncryptPassword(password)
}

func (u *User) VerifyPassword(password string) bool {
	return auth.VerifyPassword(u.PasswordDigest, password)
}

func (u *User) GenerateToken() (string, error) {
	return auth.GenerateToken(u.Jti)
}