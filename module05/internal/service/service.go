package service

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	var (
		num1 = 45
		num2 = 46
		wg   = sync.WaitGroup{}
	)

	wg.Add(2)
	go fibRun(&wg, num1)
	go fibRun(&wg, num2)
	go spinner()

	wg.Wait()
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

func fibRun(wg *sync.WaitGroup, num int) {
	res := fib(num)
	fmt.Printf("\rFibonacci(%d) = %d\n", num, res)
	wg.Done()
}
