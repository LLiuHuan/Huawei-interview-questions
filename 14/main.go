// 字符串排序

/*
题目描述
给定n个字符串，请对n个字符串按照字典序排列。
输入描述:
输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
输出描述:
数据输出n行，输出结果为按照字典序排列的字符串。
示例1

输入
9
cap
to
cat
card
two
too
up
boat
boot

输出
boat
boot
cap
card
cat
to
too
two
up
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s1, _, _ := reader.ReadLine()
	num, _ := strconv.Atoi(string(s1))
	strList := make([]string, 0)
	for i := 0; i < num; i++ {
		s2, _, _ := reader.ReadLine()
		strList = append(strList, string(s2))
	}
	sort.Strings(strList)
	for i := 0; i < len(strList); i++ {
		fmt.Println(strList[i])
	}
}
