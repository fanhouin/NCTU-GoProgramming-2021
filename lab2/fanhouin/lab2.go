package main

import "fmt"

func Sum(n int64) {
	var i int64
	var sum int64 = 1
	fmt.Print("1")
	for i = 2; i <= n; i++ {
		if i%7 != 0 {
			fmt.Printf("+%d", i)
			sum += i
		}
	}
	fmt.Printf("=%d\n", sum)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0 || n > 10000 {
			break
		}
		Sum(n)
	}
}
