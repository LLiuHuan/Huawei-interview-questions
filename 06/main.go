package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 质数因子
// 1. 如果小于2 那就需不要解了
// 2. 初始i=2 判断值取余i 如果等于0 那就除i 如果不等于 那就++
// 3. 循环结束有可能没完事 最后在判断一下 如果还有值 那就将值打印一下
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
