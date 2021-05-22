package models

import (
	"gopkg.in/guregu/null.v4"
)

type Dashboard struct {
	BaseModel
	UsersCount null.Int
}

func (Dashboard) TableName() string {
	return "dashboard"
}
