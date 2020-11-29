//errors使用
/**
fmt.Errorf在下层抛出的错误重新包装一个新错误, 但最大的问题是没有错误堆栈
 */

package main

import (
	"errors"
	"fmt"
)

type myError struct {
	Code int64
	Msg  string
}

func (err *myError) Error() string {
	return err.Msg
}

func New(code int64, text string) error {
	return &myError{code, text}
}

func main() {
	api()
}

func api() {
	err := service()

	daoErr := errors.Unwrap(err)      //返回上一级错误 output dao error
	nilError := errors.Unwrap(daoErr) //output nil
	fmt.Printf("%#v \n-  %#v \n-  %+v\n", daoErr, nilError, err)
	//output
	//   &main.myError{Code:0, Msg:"dao出错了"}
	//-  <nil>
	//-  &fmt.wrapError{msg:"service捕获dao错误 dao出错了", err:(*main.myError)(0xc00000c060)}

	fmt.Println(errors.Is(err, MY_ERROR)) //判断是否完全相等 true

	var asError *myError
	fmt.Println(errors.As(err, &asError)) //判断类型相等即可 true
}

func service() error {
	if err := dao(); err != nil {
		return fmt.Errorf("service捕获dao错误 %w", err)
	}

	return nil
}

var MY_ERROR = New(0, "dao出错了")
func dao() error {
	return MY_ERROR
}
