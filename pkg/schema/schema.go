package schema

import (
	"database/sql"
	"time"
)

type DBItem struct {
	ID        int64  // DATABASE_PK
	Name      string // NAME
	URL       string // URL
	AccountId string // ID
	AccountPw string // PW
	Port      string // PORT
	TargetDB  string // TARGET_DB
}

type FileItem struct {
	ID   int64  // FILE_PK
	Name string // NAME
	Path string // PATH
}

type History struct {
	ID             int64         // HISTORY_PK
	RunDate        time.Time     // RUN_DATE
	Status         bool          // STATUS (성공: 1, 실패: 0)
	DatabaseID     sql.NullInt64 // DATABASE_PK
	FileID         sql.NullInt64 // FILE_PK
	Type           bool          // TYPE (데이터베이스: 1, 파일: 0)
	BackupFileName string        // BACKUP_FILE_NAME
}

func (h *History) IsSuccess() bool {
	return h.Status
}

func (h *History) IsFailure() bool {
	return !h.Status
}

func (h *History) IsDatabase() bool {
	return h.Type
}

func (h *History) IsFile() bool {
	return !h.Type
}

func (h *History) GetDatabaseID() int64 {
	if h.DatabaseID.Valid {
		return h.DatabaseID.Int64
	}
	return 0
}

func (h *History) GetFileID() int64 {
	if h.FileID.Valid {
		return h.FileID.Int64
	}
	return 0
}
