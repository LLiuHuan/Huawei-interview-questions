// 字符串分隔

/**
题目描述
•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。

输入描述:
连续输入字符串(输入多次,每个字符串长度小于100)

输出描述:
输出到长度为8的新字符串数组

示例1

输入
abc
123456789

输出
abc00000
12345678
90000000
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
