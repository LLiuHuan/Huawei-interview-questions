// 取近似值

/**
题目描述
写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于5,向上取整；小于5，则向下取整。

输入描述:
输入一个正浮点数值

输出描述:
输出该数值的近似整数值

示例1

输入
5.5

输出
6
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()
	v, _ := strconv.ParseFloat(string(s), 64)
	fmt.Println(round(v))
}

func round(v float64) float64 {
	return math.Floor(v + 0.5)
}

//func main(){
//	var num float32
//	fmt.Scanf("%f",&num)
//	fmt.Println(int(num+0.5))
//}
