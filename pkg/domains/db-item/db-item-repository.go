package db_item

import (
	"database/sql"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
)

func DeleteDatabaseItemById(tx *sql.Tx, id int) error {

	query := "DELETE FROM DB_ITEM WHERE DB_ITEM_PK = ?"

	_, err := tx.Exec(query, id)

	return err
}

func ExistDatabaseItemById(tx *sql.Tx, id int) (bool, error) {
	query := "SELECT COUNT(*) FROM DB_ITEM WHERE DB_ITEM_PK = ?"

	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func ExistDatabaseItemByName(tx *sql.Tx, name string) (bool, error) {
	query := "SELECT COUNT(*) FROM DB_ITEM WHERE NAME = ?"

	var count int
	err := tx.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetDatabaseItems(tx *sql.Tx, pageable paging.Pageable) ([]schema.DBItem, error) {
	query := "SELECT DB_ITEM_PK, NAME, URL, ID, PW, PORT, TARGET_DB FROM DB_ITEM ORDER BY ID DESC LIMIT ? OFFSET ? "

	rows, err := tx.Query(query, pageable.Limit(), pageable.Offset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []schema.DBItem

	for rows.Next() {
		item := schema.DBItem{}
		err := rows.Scan(&item.ID, &item.Name, &item.URL, &item.AccountId, &item.AccountPw, &item.Port, &item.TargetDB)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func CreateDatabaseItem(tx *sql.Tx, dbItem schema.DBItem) error {

	query := "INSERT INTO DB_ITEM (NAME, URL, ID, PW, PORT, TARGET_DB) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := tx.Exec(query, dbItem.Name, dbItem.URL, dbItem.AccountId, dbItem.AccountPw, dbItem.Port, dbItem.TargetDB)

	return err
}

func GetDatabaseItemById(tx *sql.Tx, id int) (*schema.DBItem, error) {
	query := "SELECT DB_ITEM_PK, NAME, URL, ID, PW, PORT, TARGET_DB FROM DB_ITEM WHERE DB_ITEM_PK = ?"

	var item schema.DBItem
	err := tx.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.URL, &item.AccountId, &item.AccountPw, &item.Port, &item.TargetDB)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetDatabaseItemCnt(tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(*) FROM DB_ITEM"

	var count int
	err := tx.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
