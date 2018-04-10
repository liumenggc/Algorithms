package main

import "fmt"
//斐波列契函数
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		//引入外部自由变量a, b
		a, b = b, b+a
		return b - a
	}
}

func main() {
	f := Fibonacci()

	for i := 0; i < 30 ; i++  {
		fmt.Printf("%d " ,f())
	}
}