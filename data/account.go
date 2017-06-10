package data

import (
	"database/sql"

	"go-ws/entity"
)

const (
	QUERY_SELECT_ACCOUNT_BY_ID = `SELECT * FROM account WHERE id=?`
)

func (db *DB) FindAccountById(id int64) (*entity.Account, error) {
	account := &entity.Account{}
	err := db.Get(account, QUERY_SELECT_ACCOUNT_BY_ID, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}
