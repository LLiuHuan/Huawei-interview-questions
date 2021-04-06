// 字符串最后一个单词的长度

/**
题目描述
计算字符串最后一个单词的长度，单词以空格隔开。

输入描述:
输入一行，代表要计算的字符串，非空，长度小于5000。

输出描述:
输出一个整数，表示输入字符串最后一个单词的长度。

示例1

输入
hello nowcoder

输出
8
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
