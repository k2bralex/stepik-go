package service

import (
	"fmt"
	"time"
)

func Run() {
	var (
		ch   = make(chan []int)
		num1 = 45
		num2 = 46
	)

	go fibRun(ch, num1)
	go fibRun(ch, num2)
	go spinner()

	count := 0
	for v := range ch {
		fmt.Printf("\rFibonacci(%d) = %d\n", v[0], v[1])
		count++
		if count == 2 {
			close(ch)
		}
	}

	return
}

func spinner() {
	for {
		for _, r := range `-\I/` {
			fmt.Printf("\r%c", r)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibRun(c chan []int, num int) {
	res := fib(num)
	c <- []int{num, res}
}
