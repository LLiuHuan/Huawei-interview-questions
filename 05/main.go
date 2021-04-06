package main

import (
	"fmt"
)

// 进制转换

func main() {
	var a int
	for {
		_, err := fmt.Scanf("%v", &a)
		if err != nil {
			return
		}
		fmt.Printf("%d\n", a)
	}
}
