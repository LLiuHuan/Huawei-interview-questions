package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 计算某字母出现次数

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := strings.ToLower(scanner.Text())
	scanner.Scan()
	s2 := strings.ToLower(scanner.Text())

	count := 0

	for _, ch1 := range s1 {
		if string(ch1) == s2 {
			count++
		}
	}
	fmt.Println(count)
}
