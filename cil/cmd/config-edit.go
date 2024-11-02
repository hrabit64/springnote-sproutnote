package cmd

import (
	"fmt"
	"github.com/hrabit64/sproutnote/pkg/config"
	"strconv"
	"strings"
)

func RunConfigEdit(key string, value string) {

	switch strings.ToUpper(key) {
	case "DB_BACKUP_TIME":
		updateDBBackupTime(value)
		return
	case "FILE_BACKUP_TIME":
		updateFileBackupTime(value)
		return
	case "MAX_FILE_BACKUP_HISTORY":
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid MAX_FILE_BACKUP_HISTORY! Please check the number.")
			return
		}
		updateMaxFileBackupHistory(v)
		return
	case "MAX_DB_BACKUP_HISTORY":
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Invalid MAX_FILE_BACKUP_HISTORY! Please check the number.")
			return
		}
		updateMaxDbBackupHistory(v)
		return
	default:
		fmt.Println("Invalid key! Please check the key.")
	}

}

func updateDBBackupTime(value string) {
	isValid, err := config.ValidateDbBackupTime(value)
	if err != nil {
		fmt.Println("Something went wrong while validating the DB_BACKUP_TIME.")
		fmt.Println(err)
		return
	}

	if !isValid {
		fmt.Println("Invalid DB_BACKUP_TIME! Please check the time format. (HH:MM)")
		return
	}

	config.RootEnv.DbBackupTime = value
	err = config.RewriteEnv(config.RootEnv)
	if err != nil {
		fmt.Println("Something went wrong while updating the DB_BACKUP_TIME.")
		fmt.Println(err)
		return
	}

	fmt.Println("DB_BACKUP_TIME has been updated.")
	return
}

func updateFileBackupTime(value string) {
	isValid, err := config.ValidateFileBackupTime(value)
	if err != nil {
		fmt.Println("Something went wrong while validating the FILE_BACKUP_TIME.")
		fmt.Println(err)
		return
	}

	if !isValid {
		fmt.Println("Invalid FILE_BACKUP_TIME! Please check the time format. (HH:MM)")
		return
	}

	config.RootEnv.FileBackupTime = value
	err = config.RewriteEnv(config.RootEnv)
	if err != nil {
		fmt.Println("Something went wrong while updating the FILE_BACKUP_TIME.")
		fmt.Println(err)
		return
	}

	fmt.Println("FILE_BACKUP_TIME has been updated.")
	return
}

func updateMaxFileBackupHistory(value int) {
	isValid, err := config.ValidateMaxFileBackupHistory(value)
	if err != nil {
		fmt.Println("Something went wrong while validating the MAX_FILE_BACKUP_HISTORY.")
		fmt.Println(err)
		return
	}

	if !isValid {
		fmt.Println("Invalid MAX_FILE_BACKUP_HISTORY! Please check the number.")
		return
	}

	config.RootEnv.MaxFileBackupHistory = value
	err = config.RewriteEnv(config.RootEnv)
	if err != nil {
		fmt.Println("Something went wrong while updating the MAX_FILE_BACKUP_HISTORY.")
		fmt.Println(err)
		return
	}

	fmt.Println("MAX_FILE_BACKUP_HISTORY has been updated.")
	return
}

func updateMaxDbBackupHistory(value int) {
	isValid, err := config.ValidateMaxDbBackupHistory(value)
	if err != nil {
		fmt.Println("Something went wrong while validating the MAX_DB_BACKUP_HISTORY.")
		fmt.Println(err)
		return
	}

	if !isValid {
		fmt.Println("Invalid MAX_DB_BACKUP_HISTORY! Please check the number.")
		return
	}

	config.RootEnv.MaxDbBackupHistory = value
	err = config.RewriteEnv(config.RootEnv)
	if err != nil {
		fmt.Println("Something went wrong while updating the MAX_DB_BACKUP_HISTORY.")
		fmt.Println(err)
		return
	}

	fmt.Println("MAX_DB_BACKUP_HISTORY has been updated.")
	return
}
