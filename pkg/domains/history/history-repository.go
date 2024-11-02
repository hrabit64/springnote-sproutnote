package history

import (
	"database/sql"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
)

func DeleteHistoryById(tx *sql.Tx, id int) error {
	query := "DELETE FROM HISTORY WHERE HISTORY_PK = ?"
	_, err := tx.Exec(query, id)
	return err
}

func GetHistories(tx *sql.Tx, pageable paging.Pageable) ([]schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY ORDER BY HISTORY_PK DESC LIMIT ? OFFSET ?"

	rows, err := tx.Query(query, pageable.Limit(), pageable.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []schema.History

	for rows.Next() {
		item := schema.History{}
		err := rows.Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func CreateHistory(tx *sql.Tx, history *schema.History) error {
	query := "INSERT INTO HISTORY (RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.Exec(query, history.RunDate, history.Status, history.Type, history.BackupFileName, history.DatabaseID, history.FileID)
	return err
}

func GetHistoriesByDbItemId(tx *sql.Tx, dbItemId int, pageable paging.Pageable) ([]schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY WHERE DB_ITEM_PK = ? ORDER BY RUN_DATE ASC LIMIT ? OFFSET ?"

	rows, err := tx.Query(query, dbItemId, pageable.Limit(), pageable.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []schema.History

	for rows.Next() {
		item := schema.History{}
		err := rows.Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetAllHistoriesByDbItemId(tx *sql.Tx, dbItemId int) ([]schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY WHERE DB_ITEM_PK = ? ORDER BY RUN_DATE ASC"

	rows, err := tx.Query(query, dbItemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []schema.History

	for rows.Next() {
		item := schema.History{}
		err := rows.Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetHistoriesByFileItemId(tx *sql.Tx, fileItemId int, pageable paging.Pageable) ([]schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY WHERE FILE_ITEM_PK = ? ORDER BY RUN_DATE ASC LIMIT ? OFFSET ?"

	rows, err := tx.Query(query, fileItemId, pageable.Limit(), pageable.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []schema.History

	for rows.Next() {
		item := schema.History{}
		err := rows.Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func GetAllHistoriesByFileItemId(tx *sql.Tx, fileItemId int) ([]schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY WHERE FILE_ITEM_PK = ? ORDER BY RUN_DATE ASC"

	rows, err := tx.Query(query, fileItemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []schema.History

	for rows.Next() {
		item := schema.History{}
		err := rows.Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func ExistHistoryById(tx *sql.Tx, id int) (bool, error) {
	query := "SELECT COUNT(*) FROM HISTORY WHERE HISTORY_PK = ?"

	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetHistoryById(tx *sql.Tx, id int) (*schema.History, error) {
	query := "SELECT HISTORY_PK, RUN_DATE, STATUS, TYPE, BACKUP_FILE_NAME, DB_ITEM_PK, FILE_ITEM_PK FROM HISTORY WHERE HISTORY_PK = ?"

	var item schema.History
	err := tx.QueryRow(query, id).Scan(&item.ID, &item.RunDate, &item.Status, &item.Type, &item.BackupFileName, &item.DatabaseID, &item.FileID)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func ExistHistoryByBackupFileName(tx *sql.Tx, name string) (bool, error) {
	query := "SELECT COUNT(*) FROM HISTORY WHERE BACKUP_FILE_NAME = ?"

	var count int
	err := tx.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil

}
