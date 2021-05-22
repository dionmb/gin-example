package initializers

import (
	"gin_example/app"
	"gin_example/crons"
	"github.com/robfig/cron/v3"
)

func Cron()  {
	if app.Env != "development" && app.Env != "production" {
		return
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("1 * * * * *", crons.CountUsers)
	c.Start()
}