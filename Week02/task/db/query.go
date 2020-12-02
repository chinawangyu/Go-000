package db

import (
	"database/sql"
)

//模拟查询数据条数
func SearchDbCount(id int) (int, error) {
	//todo
	return 0, sql.ErrNoRows
}
