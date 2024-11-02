package cmd

import (
	"fmt"
	fileItemService "github.com/hrabit64/sproutnote/pkg/service/file-item"
)

func RunFileRm(id int) {
	isExist, err := fileItemService.ExistFileItemById(id)
	if err != nil {
		handleFileRmErr(err)
		return
	}

	if !isExist {
		fmt.Println("No file item.")
		return
	}

	isDeleted, err := fileItemService.RemoveFileItem(id)
	if err != nil {
		handleFileRmErr(err)
		return
	}

	if !isDeleted {
		fmt.Println("Failed to remove file item.")
		return
	}

	fmt.Println("File item removed.")
}

func handleFileRmErr(err error) {
	fmt.Println("Failed to remove file item.")
	fmt.Println(err)
}
