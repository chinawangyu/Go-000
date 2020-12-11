package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_jsonInt64(t *testing.T) {

	type Data struct {
		Name     string `json:"name"`
		Money_64 int64  `json:"money_64"`
		Money_32 int32  `json:"money_32"`
	}

	data := Data{
		Name:     "小明",
		Money_64: 6673221165400640161,
		Money_32: 21,
	}

	res, _ := json.Marshal(data)
	fmt.Println(string(res))

	var d2 Data
	json.Unmarshal(res, &d2)
	fmt.Printf("%+v", d2)
}
