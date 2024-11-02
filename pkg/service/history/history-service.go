package history

import (
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/database"
	dbItemRepository "github.com/hrabit64/sproutnote/pkg/domains/db-item"
	fileItemRepository "github.com/hrabit64/sproutnote/pkg/domains/file-item"
	historyRepository "github.com/hrabit64/sproutnote/pkg/domains/history"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"path/filepath"
	"strconv"
)

func ReadHistories(pageable paging.Pageable) ([]schema.History, error) {

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

	item, err := historyRepository.GetHistories(tx, pageable)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func CreateHistory(history schema.History) error {

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

	err = historyRepository.CreateHistory(tx, &history)
	if err != nil {
		return err
	}

	return nil
}

func ReadHistoriesByDbItemId(dbItemId int, pageable paging.Pageable) ([]schema.History, error) {

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

	isExist, err := dbItemRepository.ExistDatabaseItemById(tx, dbItemId)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, &serviceError.ItemNotFound{Item: strconv.Itoa(dbItemId)}
	}

	item, err := historyRepository.GetHistoriesByDbItemId(tx, dbItemId, pageable)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func ReadAllHistoriesByDbItemId(dbItemId int) ([]schema.History, error) {

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

	isExist, err := dbItemRepository.ExistDatabaseItemById(tx, dbItemId)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, &serviceError.ItemNotFound{Item: strconv.Itoa(dbItemId)}
	}

	item, err := historyRepository.GetAllHistoriesByDbItemId(tx, dbItemId)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func ReadHistoriesByFileItemId(fileItemId int, pageable paging.Pageable) ([]schema.History, error) {

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

	isExist, err := fileItemRepository.ExistFileItemById(tx, fileItemId)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, &serviceError.ItemNotFound{Item: strconv.Itoa(fileItemId)}
	}

	item, err := historyRepository.GetHistoriesByFileItemId(tx, fileItemId, pageable)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func ReadAllHistoriesByFileItemId(fileItemId int) ([]schema.History, error) {

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

	isExist, err := fileItemRepository.ExistFileItemById(tx, fileItemId)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, &serviceError.ItemNotFound{Item: strconv.Itoa(fileItemId)}
	}

	item, err := historyRepository.GetAllHistoriesByFileItemId(tx, fileItemId)

	return item, err
}

func RemoveHistoryById(id int) (bool, error) {

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

	history, err := historyRepository.GetHistoryById(tx, id)
	if err != nil {
		return false, err
	}

	err = historyRepository.DeleteHistoryById(tx, id)
	if err != nil {
		return false, err
	}

	path := filepath.Join(config.RootEnv.BackupPath, history.BackupFileName)
	result, err := utils.RemoveFile(path)
	if err != nil {
		return false, err
	}

	if !result {
		return false, nil
	}

	return true, nil
}

func ExistHistoryByBackupFileName(backupFileName string) (bool, error) {

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

	item, err := historyRepository.ExistHistoryByBackupFileName(tx, backupFileName)
	if err != nil {
		return false, err
	}

	return item, nil
}
