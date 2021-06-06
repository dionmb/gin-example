package vo

import (
	"time"
)

type TokenRes struct {
	Token string
}

type ProfileRes struct {
	Username string
}

type RepoListRes struct {
	TotalCount int64
	Records []RepoInfoRes
}

type RepoInfoRes struct {
	ID int64
	Owner string
	Name string
	Type string
	StargazersCount int64
	LastCommitAt time.Time
}