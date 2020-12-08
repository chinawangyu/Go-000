package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"task/db"
)

func NewErrNoRows() error {
	return sql.ErrNoRows
}

var NotFoundError = NewErrNoRows()

//dao层将数据真实问题返回service层
func GetDaoDataCount(id int) (int, error) {
	rows, err := db.SearchDbCount(id)

	//如果是数据库查询不到则返回NotFoundError
	if err == sql.ErrNoRows {
		//return rows, errors.Wrapf(err, "查询db count未查询到数据, param=%d", id)
		return rows, errors.Wrapf(NotFoundError, "查询db count未查询到数据, param=%d", id)
	}

	if err != nil {
		return 0, err
	}

	return rows, nil
}
