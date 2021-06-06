package main

import (
	"gin_example/initializer"
)

func main()  {
	app := initializer.Application()

	initializer.Cron()
	initializer.MachineryServer()

	app.Run()
}