package work

import (
	"github.com/hrabit64/sproutnote/pkg/core"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"time"
)

func WorkFileBackup() {
	logger, err := utils.GetLogger()
	if err != nil {
		return
	}

	defer logger.Sync()

	startTime := time.Now()

	logger.Info("File Backup >>> Start running file backup. Time: " + startTime.Format("2006-01-02 15:04:05"))

	err = core.RunAllFileBackup()
	endTime := time.Now()

	if err != nil {
		logger.Error("File Backup >>> File backup failed. Time: " + endTime.Format("2006-01-02 15:04:05") + " err: " + err.Error())
		return
	}

	logger.Info("File Backup >>> File backup success. Time: " + endTime.Format("2006-01-02 15:04:05"))

}
