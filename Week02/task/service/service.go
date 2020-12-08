package service

import (
	"errors"
	"fmt"
	"task/dao"
)

//service层根据业务场景判定error是否往上抛
func GetServiceDataCount(id int) (int, error) {
	rows, err := dao.GetDaoDataCount(id)

	//如果 err 包含sql.ErrNoRows则吞掉此错误--不建议在service层耦合db错误
	if errors.Is(err, dao.NotFoundError) {
		fmt.Println("查询不到参数")
		return 0, nil
	}

	return rows, err
}
