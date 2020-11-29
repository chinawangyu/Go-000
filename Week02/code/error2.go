//通过内部消减error!=nil判定
package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type User struct {
	Name string
	Age  uint8
	Err  error
}

func (u *User) SetName(name string) *User {
	if utf8.RuneCountInString(name) < 5 {
		u.Err = errors.New("字符串长度不能小于5")
		return u
	}

	u.Name = name
	//other code
	return u
}

func (u *User) SetAge(age uint8) *User {
	if u.Err != nil {
		//done
		return u
	}

	u.Age = age
	//other code
	return u
}
func (u *User) ToString() string {
	return u.Name + "===" + strconv.Itoa(int(u.Age))
}

func main() {

	user := &User{}
	res := user.SetName("小宇").SetAge(18).ToString()
	fmt.Println(res, user.Err) //将错误内部化, 最后再统一判定 err中 错误信息
}
