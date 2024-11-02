package mysql_utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hrabit64/sproutnote/pkg/config"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/schema"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func RunDump(item schema.DBItem) (string, error) {
	// 먼저 커넥션 상태를 점검한다.
	result, err := CheckConnection(item.URL, item.Port, item.AccountId, item.AccountPw, item.TargetDB)
	if err != nil {
		return "", err
	}

	if !result {
		return "", &serviceError.InvalidDatabaseItem{}
	}

	// 백업 디렉토리 생성
	dirName := fmt.Sprintf("db_%s_%s_%s", item.Name, time.Now().Format("20060102"), uuid.New().String())
	outputDirectory := filepath.Join(config.RootEnv.BackupPath, dirName)
	err = os.MkdirAll(outputDirectory, os.ModePerm)

	if err != nil {
		return "", err
	}

	outputFileName := filepath.Join(outputDirectory, "backup.sql")

	// mysqldump 명령어 구성
	mysqldumpCmd := []string{
		"mysqldump",
		"--single-transaction",
		"--skip-opt",
		"--extended-insert",
		"--add-drop-database",
		"--add-drop-table",
		"--no-create-db",
		"--no-create-info",
		"--default-character-set=utf8",
		"--quick",
		"--host=" + item.URL,
		"--user=" + item.AccountId,
		"--password=" + item.AccountPw,
		"--port=" + item.Port,
		item.TargetDB,
	}

	// mysqldump 실행 및 결과 파일에 쓰기
	cmd := exec.Command(mysqldumpCmd[0], mysqldumpCmd[1:]...)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return dirName, nil
}
