package cmd

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	paging "github.com/hrabit64/sproutnote/pkg/paging"
	dbItemService "github.com/hrabit64/sproutnote/pkg/service/db-item"
	"strconv"
)

const (
	DEFAULT_PAGE_SIZE = 20
)

func RunDBShow(page int) {

	cnt, err := dbItemService.ReadDatabaseItemCnt()
	if err != nil {
		handleDBShowErr(err)
		return
	}

	if cnt == 0 {
		fmt.Println("No database items.")
		return
	}

	if page > cnt/DEFAULT_PAGE_SIZE+1 {
		fmt.Println("Invalid page number.")
		return
	}
	pageable := paging.Pageable{
		Page:     page,
		PageSize: DEFAULT_PAGE_SIZE,
	}

	items, err := dbItemService.ReadDatabaseItems(pageable)
	if err != nil {
		handleDBShowErr(err)
		return
	}

	fmt.Println("<Database items> page" + strconv.Itoa(page+1) + "/" + strconv.Itoa(cnt/DEFAULT_PAGE_SIZE+1))

	var dbItemList []DbItemTableItem
	for _, item := range items {
		dbItemList = append(dbItemList, DbItemTableItem{
			Id:        strconv.Itoa(int(item.ID)),
			Name:      item.Name,
			Url:       item.URL + ":" + item.Port,
			AccountId: item.AccountId,
			TargetDB:  item.TargetDB,
		})
	}

	ui.PrintTable(dbItemList)

}

type DbItemTableItem struct {
	Id        string `header:"ID"`
	Name      string `header:"NAME"`
	Url       string `header:"URL"`
	AccountId string `header:"ACCOUNT ID"`
	TargetDB  string `header:"TARGET DB"`
}

func handleDBShowErr(err error) {

	fmt.Println("Failed to create database item.")
	fmt.Println(err)

}
