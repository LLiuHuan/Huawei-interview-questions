package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 字符串分隔

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		v := scanner.Text()
		cnt := 0
		str := strings.Builder{}
		for i := 0; i < len(v); i++ {
			cnt++
			str.WriteByte(v[i])
			if cnt == 8 {
				cnt = 0
				fmt.Println(str.String())
				str = strings.Builder{}
			}
		}
		if cnt > 0 {
			for i := cnt; i < 8; i++ {
				str.WriteByte('0')
			}
			fmt.Println(str.String())
		}
	}
}
