package db_item

import (
	"github.com/hrabit64/sproutnote/pkg/database"
	dbItemRepository "github.com/hrabit64/sproutnote/pkg/domains/db-item"
	serviceError "github.com/hrabit64/sproutnote/pkg/error/service"
	"github.com/hrabit64/sproutnote/pkg/paging"
	"github.com/hrabit64/sproutnote/pkg/schema"
)

func ReadDatabaseItems(pageable paging.Pageable) ([]schema.DBItem, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	item, err := dbItemRepository.GetDatabaseItems(tx, pageable)
	if err != nil {
		return nil, err
	}

	////db connection string 이나 계정 정보는 암호화하여 저장
	//aesSecret := config.RootEnv.DbItemSecret
	//
	//for _, targetItem := range item {
	//	targetItem.URL, err = utils.AesDecrypt(targetItem.URL, aesSecret)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	targetItem.AccountId, err = utils.AesDecrypt(targetItem.AccountId, aesSecret)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	targetItem.AccountPw, err = utils.AesDecrypt(targetItem.AccountPw, aesSecret)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	return item, nil
}

func CreateDatabaseItem(dbItem schema.DBItem) error {

	conn, err := database.GetConnect()
	if err != nil {
		return err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 이미 존재하는 이름인지 확인
	isExist, err := dbItemRepository.ExistDatabaseItemByName(tx, dbItem.Name)
	if err != nil {
		return err
	}

	if isExist {
		return &serviceError.ItemAlreadyExists{Item: dbItem.Name}
	}

	////db connection string 이나 계정 정보는 암호화하여 저장
	//aesSecret := config.RootEnv.DbItemSecret
	//
	//dbItem.URL, err = utils.AesEncrypt(dbItem.URL, aesSecret)
	//if err != nil {
	//	return err
	//}
	//
	//dbItem.AccountId, err = utils.AesEncrypt(dbItem.AccountId, aesSecret)
	//if err != nil {
	//	return err
	//}
	//
	//dbItem.AccountPw, err = utils.AesEncrypt(dbItem.AccountPw, aesSecret)
	//if err != nil {
	//	return err
	//}

	err = dbItemRepository.CreateDatabaseItem(tx, dbItem)
	if err != nil {
		return err
	}

	return nil
}

func RemoveDatabaseItemById(id int) (bool, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	isExist, err := dbItemRepository.ExistDatabaseItemById(tx, id)
	if err != nil {
		return false, err
	}

	if !isExist {
		return false, nil
	}

	err = dbItemRepository.DeleteDatabaseItemById(tx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadDatabaseItemById(id int) (schema.DBItem, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return schema.DBItem{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return schema.DBItem{}, err
	}

	defer tx.Rollback()

	item, err := dbItemRepository.GetDatabaseItemById(tx, id)
	if err != nil {
		return schema.DBItem{}, err
	}

	////db connection string 이나 계정 정보는 암호화하여 저장
	//aesSecret := config.RootEnv.DbItemSecret
	//
	//item.URL, err = utils.AesDecrypt(item.URL, aesSecret)
	//if err != nil {
	//	return schema.DBItem{}, err
	//}
	//
	//item.AccountId, err = utils.AesDecrypt(item.AccountId, aesSecret)
	//if err != nil {
	//	return schema.DBItem{}, err
	//}
	//
	//item.AccountPw, err = utils.AesDecrypt(item.AccountPw, aesSecret)
	//if err != nil {
	//	return schema.DBItem{}, err
	//}

	return *item, nil
}

func ReadDatabaseItemCnt() (int, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	cnt, err := dbItemRepository.GetDatabaseItemCnt(tx)

	return cnt, err
}

func ExistDatabaseItemById(id int) (bool, error) {

	conn, err := database.GetConnect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return false, err
	}

	defer tx.Rollback()

	isExist, err := dbItemRepository.ExistDatabaseItemById(tx, id)

	return isExist, err
}
