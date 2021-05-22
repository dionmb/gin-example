package main

import "gin_example/initializers"

func main()  {
	initializers.Application()

	worker := initializers.MachineryWorker()

	worker.Launch()
}
