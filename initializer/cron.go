package initializer

import (
	"gin_example/app"
	"gin_example/job/cron_job"
	"github.com/robfig/cron/v3"
)

func Cron()  {
	if app.Env != "development" && app.Env != "production" {
		return
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("1 * * * * *", cron_job.CountUsers)
	c.Start()
}