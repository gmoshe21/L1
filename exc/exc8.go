package main

import (
	"fmt"
	"strconv"
)

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func main() {
	var num int64
	var ind int
	var bit bool

	fmt.Println("Число")
	fmt.Scan(&num)
	fmt.Println("Номер бита")
	fmt.Scan(&ind)
	fmt.Println("Бит в 1 или 0")
	fmt.Scan(&bit)

	tmp := strconv.FormatInt(num, 2)
	fmt.Println("Число в двоичном коде:", tmp)

	if 64 - ind + 1 > len(tmp) {
		x := (64 - len(tmp)) - ind
		for i := 0; i < x; i++ {
			tmp = "0" + tmp
		}
		if bit == true {
			tmp = "1" + tmp
		}
	} else {
		if bit == true {
			tmp = replaceAtIndex(tmp, '1', len(tmp) - (64 - ind + 1))
		} else {
			tmp = replaceAtIndex(tmp, '0', len(tmp) - (64 - ind + 1))
		}
	}

	fmt.Println("Преобразование:", tmp)
	i, err := strconv.ParseInt(tmp, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

