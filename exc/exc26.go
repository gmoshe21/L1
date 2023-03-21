package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	set := make(map[rune]struct{})

	rns := []rune(strings.ToLower(s))
	for _, rn := range rns {
		_, ok := set[rn]; if !ok {
			set[rn] = struct{}{}
		} else {
			return false
		}

	}
	return true
}

func main() {
	var str string

	fmt.Scan(&str)

	data := make(map[byte]int)

	str = (strings.ToLower(str))
	for i := 0; i < len(str); i++ {
		check := data[str[i]]; 
		if check == 0 {
			data[str[i]] = 1
		} else {
			fmt.Println("false")
			return
		}

	}
	fmt.Println("true")
}