package file_item

import (
	"github.com/hrabit64/sproutnote/pkg/database"
	fileItemRepository "github.com/hrabit64/sproutnote/pkg/domains/file-item"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
)

func ReadFileItems(pageable paging.Pageable) ([]schema.FileItem, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	item, err := fileItemRepository.GetFileItems(tx, pageable)
	if err != nil {
		return nil, err
	}

	return item, nil

}

func CreateFileItem(fileItem schema.FileItem) error {

	conn, err := database.GetConnect()
	if err != nil {
		return err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	isExist, err := fileItemRepository.ExistFileItemByName(tx, fileItem.Name)
	if err != nil {
		return err
	}

	if isExist {
		return &serviceError.ItemAlreadyExists{Item: fileItem.Name}
	}

	err = fileItemRepository.CreateFileItem(tx, fileItem)

	return err
}

func RemoveFileItem(id int) (bool, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	isExist, err := fileItemRepository.ExistFileItemById(tx, id)
	if err != nil {
		return false, err
	}

	if !isExist {
		return false, nil
	}

	err = fileItemRepository.DeleteFileItemById(tx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadFileItemById(id int) (schema.FileItem, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return schema.FileItem{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return schema.FileItem{}, err
	}

	defer tx.Rollback()

	item, err := fileItemRepository.GetFileItemById(tx, id)

	return item, err
}

func ReadFileItemCnt() (int, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	cnt, err := fileItemRepository.GetFileItemCnt(tx)

	return cnt, err
}

func ExistFileItemById(id int) (bool, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return false, err
	}

	defer tx.Rollback()

	isExist, err := fileItemRepository.ExistFileItemById(tx, id)

	return isExist, err
}
