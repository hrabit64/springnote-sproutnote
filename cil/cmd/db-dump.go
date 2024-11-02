package cmd

import (
	"errors"
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/core"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/utils"
)

func RunDBDumpSingle(id int) {
	fmt.Println("Warring! Do you want to dump the database? It may occur large load on the database system.")
	answer := ui.YesNo()

	if !answer {
		fmt.Println("Canceled.")
		return
	}

	err := core.RunSingleDBBackup(id)
	if err != nil {
		handleDBDumpErr(err)
		return
	}

}

func RunDBDumpAll() {
	fmt.Println("Warring! Do you want to dump the database? It may occur large load on the database system.")
	answer := ui.YesNo()

	if !answer {
		fmt.Println("Canceled.")
		return
	}

	err := core.RunAllDBBackup()
	if err != nil {
		handleDBDumpErr(err)
		return
	}
}

func handleDBDumpErr(err error) {

	logger, _ := utils.GetLogger()

	var invalidDatabaseItem *serviceError.InvalidDatabaseItem

	if errors.As(err, &invalidDatabaseItem) {
		fmt.Println("Database Connect fail.")
		return
	}
	fmt.Println("An error occurred while dumping the database.")
	logger.Error("An error occurred while dumping the database." + err.Error())
	fmt.Println(err)

}
