package main

import (
	"fmt"
)

func check(intrf interface{}) string {
	switch intrf.(type) {
	case int:
		return "Int"
	case string:
		return "String"
	case bool:
		return "Bool"
	case chan int:
		return "Channel"
	default:
		return ""
	}
}

func main() {
	var exc1 int = 10
	var exc2 string = "string"
	var exc3 bool = true
	exc4 := make(chan int)

	fmt.Println(check(exc1))
	fmt.Println(check(exc2))
	fmt.Println(check(exc3))
	fmt.Println(check(exc4))
}
