package database

import (
	"database/sql"
)

func InitSchema(conn *sql.DB) error {

	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS DB_ITEM (
			DB_ITEM_PK INTEGER PRIMARY KEY AUTOINCREMENT,
			NAME TEXT UNIQUE,
			URL TEXT,
			PORT TEXT,
			ID TEXT,
			PW TEXT,
			TARGET_DB TEXT
		);
		
		CREATE TABLE IF NOT EXISTS FILE_ITEM (
			FILE_ITEM_PK INTEGER PRIMARY KEY AUTOINCREMENT,
			NAME TEXT UNIQUE,
			PATH TEXT UNIQUE
		);
		
		CREATE TABLE IF NOT EXISTS HISTORY (
			HISTORY_PK INTEGER PRIMARY KEY AUTOINCREMENT,
			RUN_DATE TIMESTAMP,
			STATUS BOOLEAN,
			TYPE BOOLEAN,
			BACKUP_FILE_NAME TEXT UNIQUE,
			DB_ITEM_PK INTEGER,
			FILE_ITEM_PK INTEGER,
			FOREIGN KEY (DB_ITEM_PK) REFERENCES DB_ITEM(DB_ITEM_PK),
			FOREIGN KEY (FILE_ITEM_PK) REFERENCES FILE_ITEM(FILE_ITEM_PK)
		);
	`)

	if err != nil {
		return err
	}

	return nil

}

func RunSetup() error {

	conn, err := GetConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = InitSchema(conn)
	if err != nil {
		return err
	}

	return nil

}
