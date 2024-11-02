package cmd

import (
	"fmt"
	dbItemService "github.com/hrabit64/sproutnote/pkg/service/db-item"
)

func RunDBRemove(id int) {
	// Check if the database item exists
	isExist, err := dbItemService.ExistDatabaseItemById(id)
	if err != nil {
		handleDBRemoveErr(err)
		return
	}

	if !isExist {
		fmt.Println("Database item does not exist.")
		return
	}

	// Remove the database item
	deleted, err := dbItemService.RemoveDatabaseItemById(id)
	if err != nil {
		handleDBRemoveErr(err)
		return
	}

	if deleted {
		fmt.Println("Database item removed successfully.")
	} else {
		fmt.Println("Database item could not be removed.")
	}
}

func handleDBRemoveErr(err error) {
	fmt.Println("Error occurred while removing the database item.")
}
