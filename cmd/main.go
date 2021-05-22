package main

import (
	"gin_example/initializers"
)

func main()  {
	app := initializers.Application()

	initializers.Cron()
	initializers.MachineryServer()

	app.Run()
}