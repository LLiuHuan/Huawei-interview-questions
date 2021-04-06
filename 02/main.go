// 计算某字母出现次数

/**
题目描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写。

输入描述:
第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母。

输出描述:
输出输入字符串中含有该字符的个数。

示例1

输入
ABCabc
A

输出
2
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
