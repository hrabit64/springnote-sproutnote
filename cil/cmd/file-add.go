package cmd

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/schema"
	fileItemService "github.com/hrabit64/sproutnote/pkg/service/file-item"
	"github.com/hrabit64/sproutnote/pkg/utils"
)

func RunFileAdd() {
	fmt.Println("1. Enter the file item display name. (1 to 50 characters) : ")
	name, err := ui.InputBox("Display Name")
	if err != nil {
		handleFileAddErr(err)
		return
	}
	if len(name) > 50 || len(name) == 0 {
		fmt.Println("Invalid file item name. Please enter a name with 1 to 50 characters.")
		return
	}
	fmt.Println("2. Enter the file item path : ")
	path, err := ui.InputBox("Path")
	if err != nil {
		handleFileAddErr(err)
		return
	}
	if len(path) == 0 {
		fmt.Println("Invalid file item path.")
		return
	}
	fmt.Println("3. Confirm the file item information.")
	fmt.Println("Name : " + name)
	fmt.Println("Path : " + path)
	fmt.Println("Do you want to add this file item? (y/n)")

	confirm := ui.YesNo()
	if !confirm {
		fmt.Println("Canceled.")
		return
	}

	isExist, err := utils.CheckFileExist(path)
	if err != nil {
		handleFileAddErr(err)
		return
	}
	if !isExist {
		fmt.Println("File or Dir is not exist.")
		return
	}

	fileItem := schema.FileItem{
		Name: name,
		Path: path,
	}

	err = fileItemService.CreateFileItem(fileItem)
	if err != nil {
		handleFileAddErr(err)
		return
	}

	fmt.Println("File item added successfully.")
}

func handleFileAddErr(err error) {
	fmt.Println("Failed to add file item.")
	fmt.Println(err)
}
