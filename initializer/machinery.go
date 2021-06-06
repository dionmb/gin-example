package initializer

import (
	"gin_example/app"
	"gin_example/job/delay_job"
	"gin_example/lib/configuration"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"log"
)

type machineryConfig struct {
	Brokers []string
	Backends []string
}

func loadMachineryConfig() machineryConfig {
	var config machineryConfig
	configuration.LoadConfig("machinery", &config)
	return config
}


func registerTasks(server * machinery.Server)  {
	tasks := map[string]interface{}{
	"CountUsers" : delay_job.CountUsers,
	}

	if err :=  server.RegisterTasks(tasks); err != nil {
		log.Fatalln(err)
	}
}

func MachineryServer() *machinery.Server {
	if app.Machinery != nil {
		return app.Machinery
	}

	conf := loadMachineryConfig()

	cnf := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	broker := redisbroker.NewGR(cnf, conf.Brokers, 0)
	backend := redisbackend.NewGR(cnf, conf.Backends, 0)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	registerTasks(server)

	app.Machinery = server
	return server
}

func MachineryWorker() *machinery.Worker {
	consumerTag := "machinery_worker"
	server := MachineryServer()

	worker := server.NewWorker(consumerTag, 0)

	errorhandler := func(err error) {
		log.Println("worker error:", err)
	}

	worker.SetErrorHandler(errorhandler)

	return worker
}