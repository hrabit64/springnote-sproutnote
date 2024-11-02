package cmd

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/core"
)

func RunFileBackupAll() {
	fmt.Println("Warring! Do you want to backup the file? It may occur large load on the system.")
	answer := ui.YesNo()

	if !answer {
		fmt.Println("Canceled.")
		return
	}

	fmt.Println("Backing up all files...")

	err := core.RunAllFileBackup()
	if err != nil {
		handleFileBackupErr(err)
		return
	}
}

func RunFileBackupSingle(id int) {
	fmt.Println("Warring! Do you want to backup the file? It may occur large load on the system.")
	answer := ui.YesNo()

	if !answer {
		fmt.Println("Canceled.")
		return
	}

	err := core.RunSingleFileBackup(id)
	if err != nil {
		handleFileBackupErr(err)
		return
	}
}

func handleFileBackupErr(err error) {
	fmt.Println("An error occurred while backing up the file.")
	fmt.Println(err)
}
