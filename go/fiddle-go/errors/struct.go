package main

import (
	"fmt"
	"os"
)

// エラー処理用の構造体
type MyError struct {
	Msg  string
	Code int
}

// MyError構造体にerrorインターフェースのError関数を実装
func (err *MyError) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.Msg, err.Code)
}

func main() {
	if err := doError(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("the end")
}

func doError() error {
	return &MyError{Msg: "nothig at all", Code: 19}
}
