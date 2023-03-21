package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()

	str := strings.Split(myscanner.Text(), " ")
	for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
		str[i], str[j] = str[j], str[i]
	}

	fmt.Println(strings.Join(str, " "))
}