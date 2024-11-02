package work

import (
	"github.com/hrabit64/sproutnote/pkg/core"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"time"
)

func WorkDbDump() {
	logger, err := utils.GetLogger()
	if err != nil {
		return
	}
	defer logger.Sync()

	startTime := time.Now()

	logger.Info("DB Dump  >>> Start running DB dump. Time: " + startTime.Format("2006-01-02 15:04:05"))

	err = core.RunAllDBBackup()
	endTime := time.Now()

	if err != nil {
		logger.Error("DB Dump >>> DB dump failed. Time: " + endTime.Format("2006-01-02 15:04:05") + " err: " + err.Error())
		return
	}

	logger.Info("DB Dump  >>> DB dump success. Time: " + endTime.Format("2006-01-02 15:04:05"))

	return
}
