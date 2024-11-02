package file_item

import (
	"database/sql"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
)

func DeleteFileItemById(tx *sql.Tx, id int) error {
	query := "DELETE FROM FILE_ITEM WHERE id = ?"
	_, err := tx.Exec(query, id)
	return err
}

func ExistFileItemById(tx *sql.Tx, id int) (bool, error) {
	query := "SELECT COUNT(*) FROM FILE_ITEM WHERE FILE_ITEM_PK = ?"

	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func ExistFileItemByName(tx *sql.Tx, name string) (bool, error) {
	query := "SELECT COUNT(*) FROM FILE_ITEM WHERE NAME = ?"

	var count int
	err := tx.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetFileItems(tx *sql.Tx, pageable paging.Pageable) ([]schema.FileItem, error) {
	query := "SELECT FILE_ITEM_PK, NAME, PATH FROM FILE_ITEM LIMIT ? OFFSET ?"

	rows, err := tx.Query(query, pageable.PageSize, pageable.Page*pageable.PageSize)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []schema.FileItem

	for rows.Next() {
		item := schema.FileItem{}
		err := rows.Scan(&item.ID, &item.Name, &item.Path)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func CreateFileItem(tx *sql.Tx, fileItem schema.FileItem) error {
	query := "INSERT INTO FILE_ITEM (NAME, PATH) VALUES (?, ?)"
	_, err := tx.Exec(query, fileItem.Name, fileItem.Path)
	return err
}

func GetFileItemById(tx *sql.Tx, id int) (schema.FileItem, error) {
	query := "SELECT FILE_ITEM_PK, NAME, PATH FROM FILE_ITEM WHERE FILE_ITEM_PK = ?"

	item := schema.FileItem{}
	err := tx.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Path)
	if err != nil {
		return schema.FileItem{}, err
	}

	return item, nil
}

func GetFileItemCnt(tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(*) FROM FILE_ITEM"

	var count int
	err := tx.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
