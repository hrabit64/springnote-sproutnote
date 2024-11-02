package cmd

import (
	"fmt"
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/paging"
	fileItemService "github.com/hrabit64/sproutnote/pkg/service/file-item"
	"strconv"
)

func RunFileShow(page int) {
	pageable := paging.Pageable{
		Page:     page,
		PageSize: DEFAULT_PAGE_SIZE,
	}

	cnt, err := fileItemService.ReadFileItemCnt()
	if cnt < 1 {
		fmt.Println("No file items.")
		return
	}

	items, err := fileItemService.ReadFileItems(pageable)
	if err != nil {
		handleFileShowErr(err)
		return
	}

	fmt.Println("<File items> page" + strconv.Itoa(page+1) + "/" + strconv.Itoa(cnt/DEFAULT_PAGE_SIZE+1))

	var fileItemList []FileItemList
	for _, item := range items {
		fileItemList = append(fileItemList, FileItemList{
			Id:   strconv.Itoa(int(item.ID)),
			Name: item.Name,
			Path: item.Path,
		})
	}

	ui.PrintTable(fileItemList)
}

type FileItemList struct {
	Id   string `header:"ID"`
	Name string `header:"Name"`
	Path string `header:"Path"`
}

func handleFileShowErr(err error) {
	fmt.Println("Failed to show file items.")
	fmt.Println(err)
}
