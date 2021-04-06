// 质数因子
// 1. 如果小于2 那就需不要解了
// 2. 初始i=2 判断值取余i 如果等于0 那就除i 如果不等于 那就++
// 3. 循环结束有可能没完事 最后在判断一下 如果还有值 那就将值打印一下

/**
题目描述
功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）

最后一个数后面也要有空格

输入描述:
输入一个long型整数

输出描述:
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。

示例1

输入
180

输出
2 2 3 3 5
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		v, _, _ := reader.ReadLine()
		in, err := strconv.Atoi(string(v))

		if err != nil {
			break
		}

		if in < 2 {
			fmt.Println(in)
			continue
		}

		for i := 2; i*i <= in; {
			if in%i == 0 {
				in /= i
				fmt.Printf("%d ", i)
			} else {
				i++
			}
		}

		if in > 1 {
			fmt.Printf("%d\n", in)
		}
	}
}
