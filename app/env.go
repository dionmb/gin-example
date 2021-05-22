package app

import (
	"github.com/RichardKnop/machinery/v2"
	"gorm.io/gorm"
)

type ApplicationConfig struct {
	JwtSecret string
}

var Env string
var DB *gorm.DB
var Root string
var Machinery *machinery.Server
var Config ApplicationConfig
