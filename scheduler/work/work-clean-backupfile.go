package work

import (
	"github.com/hrabit64/sproutnote/pkg/config"
	historyService "github.com/hrabit64/sproutnote/pkg/service/history"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"os"
	"path"
)

func WorkCleanBackupFile() {
	logger, err := utils.GetLogger()
	if err != nil {
		return
	}
	defer logger.Sync()

	logger.Info("CleanFile >>> Start running clean backup file.")

	targetDir := config.RootEnv.BackupPath

	files, err := os.ReadDir(targetDir)
	if err != nil {
		logger.Error("CleanFile >>> Clean backup file failed. err: " + err.Error())
		return
	}

	for _, file := range files {
		isExist, err := historyService.ExistHistoryByBackupFileName(file.Name())
		if err != nil {
			logger.Error("CleanFile >>> Check history exist failed." + file.Name() + "err: " + err.Error())
			continue
		}

		if isExist {
			continue
		}

		ok, err := utils.RemoveFile(path.Join(targetDir, file.Name()))
		if err != nil {
			logger.Error("CleanFile >>> Remove file failed." + file.Name() + "err: " + err.Error())
		}

		if !ok {
			logger.Info("CleanFile >>> Remove file failed." + file.Name())
		} else {
			logger.Info("CleanFile >>> Remove file success." + file.Name())
		}
	}
}
