package main

import "gin_example/initializer"

func main()  {
	initializer.Application()

	worker := initializer.MachineryWorker()

	worker.Launch()
}
