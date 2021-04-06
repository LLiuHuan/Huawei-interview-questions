package main

/**
题目描述
输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是0。

输入描述:
输入一个int型整数

输出描述:
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数

示例1

输入:
9876673

输出:
37689

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	v, _, _ := reader.ReadLine()
	newMap := make(map[int]bool)
	for i := len(v) - 1; i >= 0; i-- {
		intV, _ := strconv.Atoi(string(v[i]))
		if !newMap[intV] {
			newMap[intV] = true
			fmt.Print(intV)
		}
	}
}
