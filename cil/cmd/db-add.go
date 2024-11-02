package cmd

import (
	"errors"
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/schema"
	dbItemService "github.com/hrabit64/sproutnote/pkg/service/db-item"
	mysqlutils "github.com/hrabit64/sproutnote/pkg/utils/mysql-utils"
	"github.com/hrabit64/sproutnote/pkg/utils/regex"
	"regexp"
)

func RunDBCreate() {
	fmt.Println("1. Enter Display Database Name. (1 to 50 characters, for only display. not connection url) : ")
	name, err := ui.InputBox("Display Name")
	if err != nil {
		handleDBCreateErr(err)
		return
	}
	if len(name) > 50 || len(name) == 0 {
		fmt.Println("Invalid database name. Please enter a name with 1 to 50 characters.")
		return
	}

	fmt.Println("2. Enter Database URL (do not include port) : ")
	url, err := ui.InputBox("Database URL")
	if err != nil {
		handleDBCreateErr(err)
		return
	}
	if len(url) == 0 {
		fmt.Println("Invalid database url.")
		return
	}

	fmt.Println("3. Enter Database Port : ")
	port, err := ui.InputBox("Database Port")
	if err != nil {
		handleDBCreateErr(err)
		return
	}
	result, err := regexp.MatchString(regex.PORT_REGEX, port)
	if err != nil || !result {
		fmt.Print("Invalid port number.")
		return
	}

	fmt.Println("4. Enter Database Account ID : ")
	accountId, err := ui.InputBox("Account ID")
	if err != nil {
		handleDBCreateErr(err)
		return
	}

	fmt.Println("5. Enter Database Account Password : ")
	password, err := ui.Password()
	if err != nil {
		fmt.Println("Failed to get password.")
		return
	}

	fmt.Println("6. Enter Target Database Name : ")
	targetDb, err := ui.InputBox("Target Database Name")
	if err != nil {
		handleDBCreateErr(err)
		return
	}

	fmt.Println("7. Confirm the entered information.")
	fmt.Println("Database Name : ", name)
	fmt.Println("Database URL : ", url)
	fmt.Println("Database Port : ", port)
	fmt.Println("Database Account ID : ", accountId)
	fmt.Println("Database Target DB : ", targetDb)
	fmt.Println("Is the information correct? (y/n)")

	confirm := ui.YesNo()
	if confirm {
		fmt.Println("Start Add Database Item")
	} else {
		fmt.Println("Canceled.")
		return
	}

	connStat, err := mysqlutils.CheckConnection(url, port, accountId, password, targetDb)
	if err != nil {
		handleDBCreateErr(err)
		return
	}

	if !connStat {
		fmt.Println("Failed to connect to the database. Please check the database information.")
		return
	}

	err = dbItemService.CreateDatabaseItem(schema.DBItem{
		Name:      name,
		URL:       url,
		AccountId: accountId,
		AccountPw: password,
		Port:      port,
		TargetDB:  targetDb,
	})
	if err != nil {
		handleDBCreateErr(err)
	} else {
		fmt.Println("Successfully added database item.")
	}
}

func handleDBCreateErr(err error) {
	var itemAlreadyExists *serviceError.ItemAlreadyExists

	if errors.As(err, &itemAlreadyExists) {
		fmt.Println("Database item already exists. Please check the database name or url.")
	} else {
		fmt.Println("Failed to create database item.")
		fmt.Println(err)
	}
}

//
//func RunDBCreate(status string, arg string) {
//	switch status {
//	case runningStatus.NONE:
//		fmt.Println("1. Enter Database Name 1 to 50 characters. (for only display. not connection url) : ")
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_NM)
//
//	case runningStatus.DB_CREATE_NM:
//		if len(arg) > 50 || len(arg) == 0 {
//			fmt.Println("Invalid database name. Please enter a name with 1 to 50 characters.")
//			resetDBCreate()
//			return
//		}
//		data.SetData("name", arg)
//
//		fmt.Println("2. Enter Database URL (do not include port) : ")
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_URL)
//
//	case runningStatus.DB_CREATE_URL:
//		if len(arg) == 0 {
//			fmt.Println("Invalid database url.")
//			resetDBCreate()
//			return
//		}
//
//		data.SetData("url", arg)
//
//		fmt.Println("3. Enter Database Port : ")
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_PORT)
//
//	case runningStatus.DB_CREATE_PORT:
//		result, err := regexp.MatchString(utils.PORT_REGEX, arg)
//		if err != nil || !result {
//			fmt.Print("Invalid port number.")
//			resetDBCreate()
//			return
//		}
//
//		data.SetData("port", arg)
//
//		fmt.Println("4. Enter Database Account ID : ")
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_ACCOUNT_ID)
//
//	case runningStatus.DB_CREATE_ACCOUNT_ID:
//		if len(arg) == 0 {
//			fmt.Println("Invalid account id.")
//			resetDBCreate()
//			return
//		}
//
//		data.SetData("account_id", arg)
//
//		fmt.Println("5. Enter Database Account Password : ")
//		password, err := ui.Password()
//		if err != nil {
//			fmt.Println("Failed to get password.")
//			resetDBCreate()
//			return
//		}
//		data.SetData("account_pw", password)
//		fmt.Println("6. Enter Target Database Name : ")
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_TARGET_DB)
//
//	case runningStatus.DB_CREATE_TARGET_DB:
//		data.SetData("target_db", arg)
//		fmt.Println("7. Confirm the entered information.")
//		fmt.Println("--------------------------------")
//		fmt.Println("Database Name : ", data.GetData("name"))
//		fmt.Println("Database URL : ", data.GetData("url"))
//		fmt.Println("Database Port : ", data.GetData("port"))
//		fmt.Println("Database Account ID : ", data.GetData("account_id"))
//		fmt.Println("Database Target DB : ", data.GetData("target_db"))
//		fmt.Println("--------------------------------")
//		fmt.Println("Is the information correct? (y/n)")
//
//		runningStatus.SetRunningStatus(runningStatus.DB_CREATE_CONFIRM)
//
//	case runningStatus.DB_CREATE_CONFIRM:
//		if arg == "y" || arg == "Y" || arg == "yes" || arg == "YES" {
//			fmt.Println("Start Add Database Item")
//		} else {
//			fmt.Println("Canceled.")
//			resetDBCreate()
//			return
//		}
//		connStat, err := mysqlutils.CheckConnection(data.GetData("url"), data.GetData("port"), data.GetData("account_id"), data.GetData("account_pw"), data.GetData("target_db"))
//		if err != nil {
//			handleDBCreateErr(err)
//			resetDBCreate()
//			return
//		}
//
//		if !connStat {
//			fmt.Println("Failed to connect to the database. Please check the database information.")
//			resetDBCreate()
//			return
//		}
//
//		err = dbItemService.CreateDatabaseItem(schema.DBItem{
//			Name:      data.GetData("name"),
//			URL:       data.GetData("url"),
//			AccountId: data.GetData("account_id"),
//			AccountPw: data.GetData("account_pw"),
//			Port:      data.GetData("port"),
//			TargetDB:  data.GetData("target_db"),
//		})
//		resetDBCreate()
//		if err != nil {
//			handleDBCreateErr(err)
//		} else {
//			fmt.Println("Successfully added database item.")
//		}
//
//	default:
//		resetDBCreate()
//		fmt.Println("something is wrong")
//	}
//
//}
//func resetDBCreate() {
//	data.ClearData()
//	runningStatus.SetRunningStatus(runningStatus.NONE)
//}
