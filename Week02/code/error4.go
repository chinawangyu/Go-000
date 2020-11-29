package main

import (
	"errors"
	"fmt"
)
import xerrors "github.com/pkg/errors"

func main() {

	err := api4()
	fmt.Printf("%+v", err)
}

func api4() error {
	err := service4()
	return err
}

func service4() error {
	err := dao4()
	return xerrors.WithMessage(err, "serice层返回")
}

func dao4() error {

	err := base()
	return xerrors.Wrapf(err, "dao层返回，参数=%s", "abc")
}

func base() error {
	return errors.New("数据库出现错了啦.")
}
