/**
应用层出现错误直接使用 errors.New、errors.Errorf
调用其他包的函数通常直接返回error, 不做包装, 否则可能会返回两层堆栈
底层错误(依赖第三方库、基础库) 可以使用 errors.Wrap、errors.Wrapf 包装下错误再返回

如果降级处理, error返回nil. 自行处理错误日志

基础库不能使用wrap
*/

package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func service3() error {
	if err := dao3(); err != nil {
		return errors.WithMessage(err, "params query") //添加其他上下文信息
	}
	return nil
}

func dao3() error {
	return errors.New("dao层错误")
}

func main() {

	err := service3()
	fmt.Printf("%+v", errors.Cause(err))
	fmt.Printf("%+v\n", err)  //打印堆栈详情
}
