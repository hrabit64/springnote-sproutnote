package cmd

import (
	"errors"
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/config"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/paging"
	historyService "github.com/hrabit64/sproutnote/pkg/service/history"
	"strconv"
)

func RunHistoryFileItemShow(id int) {
	pageable := paging.Pageable{Page: 1, PageSize: config.RootEnv.MaxDbBackupHistory}
	items, err := historyService.ReadHistoriesByFileItemId(id, pageable)
	if err != nil {
		handleHistoryDBItemShowErr(err)
		return
	}

	if len(items) == 0 {
		fmt.Println("No history file item.")
		return
	}

	var historyTableItems []HistoryTableItem
	for _, item := range items {
		historyTableItems = append(historyTableItems, HistoryTableItem{
			Id:   strconv.Itoa(int(item.ID)),
			Date: item.RunDate.Format("2006-01-02 15:04:05"),
			Dir:  item.BackupFileName,
		})
	}

	fmt.Println("<History file items>")
	ui.PrintTable(historyTableItems)

}

func handleHistoryFileItemShowErr(err error) {
	var itemNotFound *serviceError.ItemNotFound

	if errors.As(err, &itemNotFound) {
		fmt.Println("No file item.")
		return
	}

	fmt.Println("Failed to show history db item.")
	fmt.Println(err)
}
