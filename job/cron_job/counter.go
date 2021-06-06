package cron_job

import (
	"gin_example/app"
	"gin_example/model"
	"log"
)

func CountUsers()  {
	log.Println("Cron Counter...")
	var count int64
	app.DB.Model(model.User{}).Count(&count)
	var dashboard model.Dashboard
	app.DB.Unscoped().FirstOrCreate(&dashboard)
	dashboard.UsersCount.SetValid(count)
	app.DB.Save(&dashboard)
}