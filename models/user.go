package models

type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null"`
	PasswordDigest string
	Activated bool `gorm:"default:true"`
	Jti string `gorm:"uniqueIndex;not null;default:uuid_generate_v4()"`
}