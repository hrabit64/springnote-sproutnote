package cmd

import (
	"github.com/hrabit64/sproutnote/cil/ui"
	"github.com/hrabit64/sproutnote/pkg/config"
	"strconv"
)

type ConfigTableItem struct {
	Key   string `header:"KEY"`
	Value string `header:"VALUE"`
}

func RunShowConfig() {

	configItems := []ConfigTableItem{
		{"BACKUP_PATH", config.RootEnv.BackupPath},
		{"FILE_BACKUP_TIME", config.RootEnv.FileBackupTime},
		{"MAX_FILE_BACKUP_HISTORY", strconv.Itoa(config.RootEnv.MaxFileBackupHistory)},
		{"DB_BACKUP_TIME", config.RootEnv.DbBackupTime},
		{"MAX_DB_BACKUP_HISTORY", strconv.Itoa(config.RootEnv.MaxDbBackupHistory)},
	}

	ui.PrintTable(configItems)

}
