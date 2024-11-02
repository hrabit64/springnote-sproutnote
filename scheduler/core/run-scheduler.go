package core

import (
	"github.com/hrabit64/sproutnote/pkg/config"
	"github.com/hrabit64/sproutnote/pkg/utils"
	"github.com/hrabit64/sproutnote/scheduler/work"
	"log"
	"os"
	"sync"
	"time"
)

var (
	fileWorkRunning = false
	dbWorkRunning   = false
	hasTrash        = false
	mu              sync.Mutex
)

func RunScheduler() {
	interval := 1 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 환경 변수 감시 고루틴 실행
	go runEnvWatch()

	for {
		select {
		case <-ticker.C:
			now := time.Now().Format("15:04")

			mu.Lock()
			if config.RootEnv.FileBackupTime == now && !fileWorkRunning {
				fileWorkRunning = true
				hasTrash = true
				go runFileWorks()
			}
			mu.Unlock()

			mu.Lock()
			if config.RootEnv.DbBackupTime == now && !dbWorkRunning {
				dbWorkRunning = true
				hasTrash = true
				go runDBWorks()
			}
			mu.Unlock()

			mu.Lock()
			if config.RootEnv.DbBackupTime != now && dbWorkRunning {
				dbWorkRunning = false
			}
			mu.Unlock()

			mu.Lock()
			if config.RootEnv.FileBackupTime != now && fileWorkRunning {
				fileWorkRunning = false
			}
			mu.Unlock()

		}
	}
}

func runEnvWatch() {
	logger, err := utils.GetLogger()
	if err != nil {
		log.Println("Failed to get logger.")
		return
	}

	configFilePath := "./.env"
	var lastModTime time.Time

	for {
		fileInfo, err := os.Stat(configFilePath)
		if err != nil {
			logger.Error("Failed to get file info. err: " + err.Error())
			break
		}

		if !fileInfo.ModTime().Equal(lastModTime) {
			logger.Info("Configuration file modified, reloading...")
			lastModTime = fileInfo.ModTime()
			config.RootEnv, err = config.LoadEnv()
		}

		time.Sleep(1 * time.Second)
	}
}

func runDBWorks() {
	work.WorkDbDump()
	work.WorkCleanBackupFile()
}

func runFileWorks() {
	work.WorkFileBackup()
	work.WorkCleanBackupFile()
}
