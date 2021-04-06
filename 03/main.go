package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 明明的随机数

func main() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		in, _, _ := scanner.ReadLine()
		if len(in) == 0 {
			break
		}
		count, _ := strconv.Atoi(string(in))

		mark := [1000]bool{}
		for i := 0; i < count; i++ {
			txt, _, _ := scanner.ReadLine()
			itxt, _ := strconv.Atoi(string(txt))
			mark[itxt] = true
		}
		for k, v := range mark {
			if v {
				fmt.Println(k)
			}
		}
	}
}
