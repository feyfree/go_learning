package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"head_first_go/ch10/mytype"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date)
	newType := mytype.MyNewType{}
	newType.ExportMethod()
	// 小写不能调用的
}
