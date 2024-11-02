package core

import (
	"database/sql"
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
	dbItemService "github.com/hrabit64/sproutnote/pkg/service/db-item"
	historyService "github.com/hrabit64/sproutnote/pkg/service/history"
	"github.com/hrabit64/sproutnote/pkg/utils"
	mysqlUtils "github.com/hrabit64/sproutnote/pkg/utils/mysql-utils"
	"strconv"
	"time"
)

func runDBBackup(dbItem schema.DBItem) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	// Dump the database
	dumpFileName, err := mysqlUtils.RunDump(dbItem)
	if err != nil {

		logger.Error("MysqlDump >>> Failed to dump the database. Database name : " + dbItem.Name + " Error : " + err.Error())

		return err

	}

	logger.Info("MysqlDump >>> Database dumped successfully. File name : " + dumpFileName)

	// 백업 정보를 바탕으로 history 생성
	newHistory := schema.History{
		RunDate:        time.Now(),
		Status:         true,
		DatabaseID:     sql.NullInt64{Int64: dbItem.ID, Valid: true},
		Type:           true,
		BackupFileName: dumpFileName,
	}

	err = historyService.CreateHistory(newHistory)
	if err != nil {
		return err
	}

	histories, err := historyService.ReadAllHistoriesByDbItemId(int(dbItem.ID))
	if err != nil {
		return err
	}

	if len(histories) > config.RootEnv.MaxDbBackupHistory {
		logger.Info("MysqlDump >>> Full history count. Removing the oldest history.")

		historiesID := make([]int, len(histories))
		for i, history := range histories {
			historiesID[i] = int(history.ID)
		}

		// 오래된 히스토리만큼 삭제
		excess := len(histories) - config.RootEnv.MaxDbBackupHistory
		for i := 0; i < excess; i++ {
			logger.Info("MysqlDump >>> Removing the oldest history. ID : " + strconv.Itoa(historiesID[i]))
			err := deleteDBHistory(historiesID[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RunAllDBBackup() error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	cnt, err := dbItemService.ReadDatabaseItemCnt()
	if err != nil {
		return err
	}

	if cnt == 0 {
		logger.Info("MysqlDump >>> No database items.")
		return nil
	}

	logger.Info("MysqlDump >>> Found " + strconv.Itoa(cnt) + " database items.")

	err = dumpItems(cnt)
	if err != nil {
		return err
	}
	return err
}

func dumpItems(cnt int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

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

		items, err := dbItemService.ReadDatabaseItems(pageable)
		if err != nil {
			return err
		}

		for _, item := range items {
			total++
			logger.Info("MysqlDump >>> ( " + strconv.Itoa(total) + "/" + strconv.Itoa(cnt) + " ) Start dumping the database ( " + item.Name + " )")
			err := runDBBackup(item)
			if err != nil {
				return err
			}
		}
		page++
	}

	return nil

}

func RunSingleDBBackup(id int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()
	isExist, err := dbItemService.ExistDatabaseItemById(id)
	if err != nil {
		return err
	}

	if !isExist {
		logger.Info("MysqlDump >>> Database item not found.")
		return nil
	}

	item, err := dbItemService.ReadDatabaseItemById(id)
	if err != nil {
		return err
	}

	logger.Info("MysqlDump >>> Start dumping the database ( " + item.Name + " )")

	err = runDBBackup(item)
	if err != nil {
		return err
	}
	return nil
}

func deleteDBHistory(id int) error {
	logger, err := utils.GetLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	result, err := historyService.RemoveHistoryById(id)
	if err != nil {
		return err
	}
	if !result {
		logger.Info("MysqlDump >>> Successfully removed the oldest history but failed to remove the backup file.")
		return nil
	}
	return nil
}
