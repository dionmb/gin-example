package initializers

import (
	"fmt"
	"gin_example/libs/configurations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
	AutoMigrate bool
}

func loadDatabaseConfig() databaseConfig {
	var config databaseConfig
	configurations.LoadConfig("database", &config)
	return config
}


func Database(models ...interface{}) *gorm.DB {
	config := loadDatabaseConfig()
	format := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai"
	dns := fmt.Sprintf(format, config.Host, config.Port, config.User, config.Password, config.Dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if config.AutoMigrate {
		migrateDatabase(db, models...)
	}
	return db
}

func migrateDatabase(db *gorm.DB, models ...interface{})  {
	err := db.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
}
