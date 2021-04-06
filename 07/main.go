package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 取近似值

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
