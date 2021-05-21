package app

import (
	"gorm.io/gorm"
)

type ApplicationConfig struct {
	JwtSecret string
}

var Env string
var DB *gorm.DB
var Root string
var Config ApplicationConfig
