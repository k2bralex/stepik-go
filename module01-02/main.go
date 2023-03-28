package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	res := ""
	for _, v := range strings.Split(s, "") {
		num, _ := strconv.Atoi(v)
		res += strconv.Itoa(num * num)
	}
	fmt.Println(res)
}

func IsPassOk() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	for _, v := range []rune(s) {
		if !(unicode.IsDigit(v) || unicode.Is(unicode.Latin, v)) {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Println("Ok")
}

func readOdds() {
	in := bufio.NewReader(os.Stdin)
	for i := 0; ; i++ {
		c, _ := in.ReadByte()
		if c == 0x0A {
			break
		}
		if i%2 != 0 {
			fmt.Print(string(c))
		}
	}
}

func firstEntry(s, sub string) int {
	return strings.Index(s, sub)
}

func isPalindrome(s string) bool {
	reversed := ""
	for _, sub := range s {
		reversed = string(sub) + reversed
	}
	return s == reversed
}

type Test struct {
	On    bool
	Ammo  int
	Power int
}

func (t *Test) Shoot() bool {
	if !t.On || t.Ammo < 1 {
		return false
	}
	t.Ammo--
	return true
}

func (t *Test) RideBike() bool {
	if !t.On || t.Power < 1 {
		return false
	}
	t.Power--
	return true
}

func sumInt(r ...int) (int, int) {
	sum := 0
	for _, v := range r {
		sum += v
	}
	return len(r), sum
}

func minimumFromFour() int {
	var sl = [4]int{}
	var min = math.MaxInt
	for i, _ := range sl {
		fmt.Scan(&sl[i])
		if min > sl[i] {
			min = sl[i]
		}
	}
	return min
}

func decToBinary(n int) string {
	s := ""
	for n > 0 {
		if (n & 1) == 1 {
			s += "1"
		} else {
			s += "0"
		}
		n >>= 1
	}
	return s
}

func fibonacci(n int64) int64 {
	a := int64(0)
	b := int64(1)

	for i := int64(0); i < n; i++ {
		temp := a
		a = b
		b = temp + b
	}
	return a
}

func IndexOfFibonacci() {
	var n float64
	fmt.Scan(&n)

	fibo := 2.078087*math.Log(n) + 1.672276
	fmt.Println(math.Round(fibo))

}

func DigitalRoot(n int) int {
	if n == 0 {
		return 0
	}
	if n%9 == 0 {
		return 9
	}
	return n % 9
}

func Max(s []int) int {
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func ContainsNumbers() {
	var a, b = "564", "8954"
	//fmt.Scan(&a, &b)
	for _, r := range []rune(a) {
		if strings.ContainsRune(b, r) {
			fmt.Printf("%s ", string(r))
		}
	}
}

func CountYearsToDeposit() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	var i = 1
	for ; x >= y; i++ {
		x += x * p / 100
	}
	fmt.Println(i)
}

func test2() {
	var num int
	for fmt.Scan(&num); num < 100; fmt.Scan(&num) {
		if num < 10 {
			continue
		}
		fmt.Println(num)
	}
}

func findMaxInput() {
	var (
		n     int
		count = map[int]int{}
		max   = 0
	)
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {

		if max < n {
			max = n
		}

		if _, ok := count[n]; !ok {
			count[n] = 1
			continue
		}

		count[n]++
	}
	fmt.Println(count[max])
}

func test1() {
	var (
		a   int
		b   int
		sum = 0
	)
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b > 9 && b < 100 && b%8 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
