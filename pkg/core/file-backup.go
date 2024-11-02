package core

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
	fileItemService "github.com/hrabit64/sproutnote/pkg/service/file-item"
	historyService "github.com/hrabit64/sproutnote/pkg/service/history"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func RunSingleFileBackup(id int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}

	isExist, err := fileItemService.ExistFileItemById(id)
	if err != nil {
		return err
	}

	if !isExist {
		logger.Info("File Backup >>> No file item.")
		return nil
	}

	item, err := fileItemService.ReadFileItemById(id)
	if err != nil {
		return err
	}

	err = runFileBackup(item)
	if err != nil {
		return err
	}

	return nil
}

func RunAllFileBackup() error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}

	cnt, err := fileItemService.ReadFileItemCnt()
	if err != nil {
		return err
	}

	if cnt == 0 {
		logger.Info("File Backup >>> Not found file item.")
		return nil
	}

	logger.Info("File Backup >>> Found total ( " + strconv.Itoa(cnt) + " ) file items.")

	err = backupWithPage(cnt)
	if err != nil {
		return err
	}

	logger.Info("File Backup >>> All file backup completed.")
	return nil
}

func backupWithPage(cnt int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}

	page := 0
	maxPage := cnt / 100
	total := 0
	for {
		if page > maxPage {
			break
		}

		pageable := paging.Pageable{
			Page:     page,
			PageSize: 100,
		}

		items, err := fileItemService.ReadFileItems(pageable)
		if err != nil {
			return err
		}

		for _, item := range items {
			total++
			logger.Info("File Backup >>> ( " + strconv.Itoa(total) + "/" + strconv.Itoa(cnt) + " ) Start backing up the file item ( " + item.Name + " )")
			err := runFileBackup(item)
			if err != nil {
				return err
			}
		}
		page++
	}

	return nil
}

func deleteFileHistory(id int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}

	result, err := historyService.RemoveHistoryById(id)
	if err != nil {
		return err
	}
	if !result {
		logger.Info("File Backup >>> Successfully removed the oldest history but failed to remove the backup file.")
		return nil
	}
	return nil
}

func runFileBackup(item schema.FileItem) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}

	isExist, err := utils.CheckFileExist(item.Path)
	if err != nil {
		return err
	}

	if !isExist {
		logger.Info("File Backup >>> Not exist path (" + item.Path + ")")
		return nil
	}

	// 백업 디렉토리 생성
	dirName := fmt.Sprintf("file_%s_%s_%s", item.Name, time.Now().Format("20060102"), uuid.New().String())
	outputDirectory := filepath.Join(config.RootEnv.BackupPath, dirName)
	err = os.MkdirAll(outputDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	isDir, err := utils.CheckIsDir(item.Path)
	if err != nil {
		return err
	}

	if isDir {
		logger.Info("File Backup >>> Start backing up the directory (" + item.Path + ")")
		err = utils.CopyDir(item.Path, outputDirectory)
		if err != nil {
			return err
		}
	} else {
		logger.Info("File Backup >>> Start backing up the file (" + item.Path + ")")
		err = utils.CopyFile(item.Path, path.Join(outputDirectory, path.Base(item.Path)))
		if err != nil {
			return err
		}
	}

	logger.Info("File Backup >>> File backup completed. Backup directory : " + dirName)

	newHistory := schema.History{
		RunDate:        time.Now(),
		Status:         true,
		FileID:         sql.NullInt64{Int64: int64(item.ID), Valid: true},
		Type:           true,
		BackupFileName: dirName,
	}

	err = historyService.CreateHistory(newHistory)
	if err != nil {
		return err
	}

	histories, err := historyService.ReadAllHistoriesByFileItemId(int(item.ID))
	if err != nil {
		return err
	}

	if len(histories) > config.RootEnv.MaxFileBackupHistory {
		logger.Info("File Backup >>> Full history count. Removing the oldest history.")

		historiesID := make([]int, len(histories))
		for i, history := range histories {
			historiesID[i] = int(history.ID)
		}

		// 오래된 히스토리만큼 삭제
		excess := len(histories) - config.RootEnv.MaxFileBackupHistory
		for i := 0; i < excess; i++ {
			logger.Info("File Backup >>> Removing the oldest history. ID : " + strconv.Itoa(historiesID[i]))
			err := deleteFileHistory(historiesID[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
