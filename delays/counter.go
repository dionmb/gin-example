package delays

import (
	"gin_example/app"
	"gin_example/models"
	"log"
)

func CountUsers() error {
	log.Println("Delay Counter...")
	var count int64
	app.DB.Model(models.User{}).Count(&count)
	var dashboard models.Dashboard
	app.DB.Unscoped().FirstOrCreate(&dashboard)
	dashboard.UsersCount.SetValid(count)
	app.DB.Save(&dashboard)

	return nil
}
