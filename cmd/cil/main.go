package main

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/core"
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/database"
	"github.com/hrabit64/sproutnote/pkg/utils"
)

func main() {
	config.ProcessType = "cil"
	fmt.Println("Welcome! SproutNote Cil 🌱")
	logger, err := utils.GetLogger()
	if err != nil {
		fmt.Println("Setup >>> Logger setup failed.")
		panic(err)
		return
	}
	defer logger.Sync()

	// Run database setup
	fmt.Println("Setup >>> Running database setup.")
	err = database.RunSetup()
	if err != nil {
		fmt.Println("Setup >>> Database setup failed.")
		logger.Error("Database setup failed" + err.Error())
		panic(err)

		return
	}
	fmt.Println("Setup >>> Database setup completed.")

	// Load environment variables

	config.RootEnv, err = config.LoadEnv()
	fmt.Println("Setup >>> Loading environment variables.")
	if err != nil {
		logger.Error("Environment setup failed" + err.Error())
		fmt.Println("Setup >>> Environment setup failed.")
		panic(err)
		return
	}
	fmt.Println("Setup >>> Environment validation started.")
	isValid, err := config.ValidateEnv(config.RootEnv)
	if err != nil {
		logger.Error("Environment validation failed" + err.Error())
		fmt.Println("Setup >>> Environment validation failed.")
		panic(err)
		return
	}

	if !isValid {
		fmt.Println("Setup >>> Environment validation failed.")
		return
	}

	fmt.Println("Setup >>> Environment validation completed.")
	fmt.Println("Setup >>> SproutNote Cil setup completed.")

	fmt.Print("\n")
	fmt.Println(`
	 ▗▄▄▖▗▄▄▖ ▗▄▄▖  ▗▄▖ ▗▖ ▗▖▗▄▄▄▖▗▖  ▗▖ ▗▄▖▗▄▄▄▖▗▄▄▄▖
	▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌  █  ▐▛▚▖▐▌▐▌ ▐▌ █  ▐▌   
	 ▝▀▚▖▐▛▀▘ ▐▛▀▚▖▐▌ ▐▌▐▌ ▐▌  █  ▐▌ ▝▜▌▐▌ ▐▌ █  ▐▛▀▀▘
	▗▄▄▞▘▐▌   ▐▌ ▐▌▝▚▄▞▘▝▚▄▞▘  █  ▐▌  ▐▌▝▚▄▞▘ █  ▐▙▄▄▖
	`)
	fmt.Print("\n")
	fmt.Println("Welcome to SproutNote Cil! 🌱")
	fmt.Println("Type 'exit' to exit the cil.")
	fmt.Println("Type 'help' to see the list of commands.")
	core.RunPrompt()
}
