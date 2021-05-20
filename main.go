package main

import (
	"gin_example/initializers"
)

func main()  {
	app := initializers.Application()

	app.Run()
}