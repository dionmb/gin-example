package models

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        null.Int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}