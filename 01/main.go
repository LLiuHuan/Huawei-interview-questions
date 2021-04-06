package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 字符串最后一个单词的长度

func main() {
	reader := bufio.NewReader(os.Stdin)

	data, _, _ := reader.ReadLine()

	splitArr := strings.Fields(string(data))

	if len(splitArr) > 0 {
		fmt.Println(len(splitArr[len(splitArr)-1]))
	} else {
		fmt.Println(0)
	}
}
