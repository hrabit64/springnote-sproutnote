package main

import (
	"fmt"
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/database"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"github.com/hrabit64/sproutnote/scheduler/core"
)

func main() {
	config.ProcessType = "scheduler"

	logger, err := utils.GetLogger()
	if err != nil {
		fmt.Println("Failed to get logger.")
		panic(err)
		return
	}

	defer logger.Sync()
	err = database.RunSetup()
	if err != nil {
		logger.Panic("Scheduler Setup >>> Database setup failed. err: " + err.Error())
		panic(err)
		return
	}

	config.RootEnv, err = config.LoadEnv()
	if err != nil {
		logger.Panic("Scheduler Setup >>> Environment load failed. err: " + err.Error())
		panic(err)
		return
	}

	isValid, err := config.ValidateEnv(config.RootEnv)
	if err != nil {
		logger.Panic("Scheduler Setup >>> Environment validation failed. err: " + err.Error())
		panic(err)
		return
	}

	if !isValid {
		logger.Panic("Scheduler Setup >>> Environment validation failed. err: Invalid environment.")
		return
	}

	logger.Info("Scheduler >>> Start running.")

	core.RunScheduler()

}
