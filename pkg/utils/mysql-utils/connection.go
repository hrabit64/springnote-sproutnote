package mysql_utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CheckConnection(url string, port string, user string, password string, database string) (bool, error) {
	conn, err := sql.Open("mysql", user+":"+password+"@tcp("+url+":"+port+")/"+database)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	_, err = conn.Exec("SELECT 1")

	if err != nil {
		return false, err
	}

	return true, nil
}
