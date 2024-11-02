package core

import (
	"github.com/c-bata/go-prompt"
	"github.com/hrabit64/sproutnote/pkg/paging"
	dbItemService "github.com/hrabit64/sproutnote/pkg/service/db-item"
	fileItemService "github.com/hrabit64/sproutnote/pkg/service/file-item"
	"strconv"
	"strings"
)

var defaultSuggestions = []prompt.Suggest{
	{Text: "help", Description: "Show help"},
	{Text: "history", Description: "Working on the history"},
	{Text: "db", Description: "Working on the database-item"},
	{Text: "file", Description: "Working on the file-item"},
	{Text: "config", Description: "Working on the configuration"},
	{Text: "exit", Description: "Exit the CLI"},
}

var dbSuggestions = []prompt.Suggest{
	{Text: "add", Description: "Add a new database-item"},
	{Text: "show", Description: "Show the database-items"},
	{Text: "dump", Description: "Dump the database-items"},
	{Text: "remove", Description: "Remove the database-item"},
}

func dbItemListSuggestions() []prompt.Suggest {
	dbItems, err := dbItemService.ReadDatabaseItems(paging.Pageable{
		Page:     0,
		PageSize: 100,
	})

	if err != nil {
		return emptySuggestions
	}

	var suggestions []prompt.Suggest
	for _, dbItem := range dbItems {
		suggestions = append(suggestions, prompt.Suggest{Text: strconv.Itoa(int(dbItem.ID)), Description: dbItem.Name})
	}
	return suggestions
}

var emptySuggestions = []prompt.Suggest{}

var fileSuggestions = []prompt.Suggest{
	{Text: "add", Description: "Add a new file-item"},
	{Text: "show", Description: "Show the files-items"},
	{Text: "remove", Description: "Remove the file-item"},
	{Text: "backup", Description: "Backup the files-item"},
}

func fileItemListSuggestions() []prompt.Suggest {
	fileItem, err := fileItemService.ReadFileItems(paging.Pageable{Page: 0, PageSize: 100})

	if err != nil {
		return emptySuggestions
	}

	var suggestions []prompt.Suggest
	for _, fileItem := range fileItem {
		suggestions = append(suggestions, prompt.Suggest{Text: strconv.Itoa(int(fileItem.ID)), Description: fileItem.Name})
	}

	return suggestions
}

var historyShowSuggestions = []prompt.Suggest{
	{Text: "db", Description: "Show DB Item history"},
	{Text: "file", Description: "Show File Item history"},
}

var helpSuggestions = []prompt.Suggest{
	{Text: "db", Description: "About the DB Item"},
	{Text: "file", Description: "About the File Item"},
	{Text: "history", Description: "About the History"},
}

var configSuggestions = []prompt.Suggest{
	{Text: "show", Description: "Show the configuration"},
	{Text: "edit", Description: "edit the configuration"},
}

var configKeySuggestions = []prompt.Suggest{
	{Text: "db_backup_time", Description: "DB backup time"},
	{Text: "file_backup_time", Description: "File backup time"},
	{Text: "max_db_backup_history", Description: "Max DB backup history"},
	{Text: "max_file_backup_history", Description: "Max File backup history"},
}

func completer(d prompt.Document) []prompt.Suggest {
	totalText := d.TextBeforeCursor()
	cursorCmd := d.GetWordBeforeCursorWithSpace()
	cursorCmd = strings.TrimSpace(cursorCmd)
	args := strings.Split(totalText, " ")
	argCount := len(args)
	switch argCount {
	case 1:
		return prompt.FilterHasPrefix(defaultSuggestions, cursorCmd, true)
	case 2:
		switch args[0] {
		case "db":
			if cursorCmd == "db" {
				return dbSuggestions
			}
			return prompt.FilterHasPrefix(dbSuggestions, cursorCmd, true)
		case "file":
			if cursorCmd == "file" {
				return fileSuggestions
			}
			return prompt.FilterHasPrefix(fileSuggestions, cursorCmd, true)
		case "history":
			if cursorCmd == "history" {
				return historyShowSuggestions
			}
			return prompt.FilterHasPrefix(historyShowSuggestions, cursorCmd, true)
		case "config":
			if cursorCmd == "config" {
				return configSuggestions
			}
			return prompt.FilterHasPrefix(configSuggestions, cursorCmd, true)
		}
	case 3:
		switch args[0] {
		case "db":
			if args[1] == "remove" || args[1] == "dump" {
				return dbItemListSuggestions()
			}
		case "file":
			if args[1] == "remove" || args[1] == "backup" {
				return fileItemListSuggestions()
			}
		case "history":

			if args[1] == "db" {
				return dbItemListSuggestions()
			}

			if args[1] == "file" {
				return fileItemListSuggestions()
			}

		case "config":

			if args[1] == "edit" {
				if cursorCmd == "edit" {
					return configKeySuggestions
				}
				return prompt.FilterHasPrefix(configKeySuggestions, cursorCmd, true)
			}
		}

	}

	return prompt.FilterHasPrefix(emptySuggestions, cursorCmd, true)

}
