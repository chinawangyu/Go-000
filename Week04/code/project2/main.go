package main

import (
	"fmt"
	"time"
)

type DialOption struct {
	f func(options *dialOptions)
}

type dialOptions struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	db           string
	password     string
}

func DialReadTimeout(d time.Duration) DialOption {
	return DialOption{func(options *dialOptions) {
		options.readTimeout = d
	}}
}

func main() {

	obj := dialOptions{}
	fmt.Printf("%+v", obj) //{re	adTimeout:0 writeTimeout:0 db: password:}

	do := dialOptions{}
	DialReadTimeout(2).f(&do)
	fmt.Printf("%+v", do) //{readTimeout:2 writeTimeout:0 db: password:}

}
