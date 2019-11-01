package main

import (
	"fmt"
)
var (
	channel = make(chan int,20)
	myres = make(map[int]int, 20)
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myres[n] = res

}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
		channel <- i
	}

	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}

}