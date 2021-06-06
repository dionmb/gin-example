package model

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Repo struct {
	BaseModel
	Owner string
	Name string
	Type string
	StargazersCount null.Int
	LastCommitAt time.Time
}