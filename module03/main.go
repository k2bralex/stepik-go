package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

const N = 5

func main() {
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		fmt.Println("getting res ", i)
		c := <-out
		fmt.Println("received res", i)
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	go func() {
		wg := &sync.WaitGroup{}
		out1 := []int{}
		out2 := []int{}

		wg.Add(2 * n)

		for i := 0; i < n; i++ {
			out1 = append(out1, <-in1)
			out2 = append(out2, <-in2)
		}

		for i := 0; i < n; i++ {
			go func(index, value int, sl *[]int) {
				defer wg.Done()
				(*sl)[index] = fn(value)
			}(i, out1[i], &out1)

			go func(index, value int, sl *[]int) {
				defer wg.Done()
				(*sl)[index] = fn(value)
			}(i, out2[i], &out2)
		}

		go func(sl1, sl2 *[]int, w *sync.WaitGroup) {

			w.Wait()

			for i := 0; i < n; i++ {
				out <- (*sl1)[i] + (*sl2)[i]
			}

			close(out)
		}(&out1, &out2, wg)

	}()

}

func calculator1(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		sum := 0
		for {
			select {
			case <-done:
				out <- sum
				return
			case val := <-arguments:
				sum += val
			}
		}
	}()

	return out
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		select {
		case <-stopChan:
		case val := <-firstChan:
			out <- val * val
		case val := <-secondChan:
			out <- val * 3
		}
	}()

	return out
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			work1()
			wg.Done()
		}()
	}

	wg.Wait()
}

func SyncGourutinesPattern() {
	<-WrapWork(work1)
}

func WrapWork(f func()) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		work1()
		close(ch)
	}()
	return ch
}

func work1() {

}

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	lib := map[string]int{}
	for in := range inputStream {
		if _, ok := lib[in]; !ok {
			lib[in] = 1
			outputStream <- in
			continue
		}
	}
}

func TimePlusDuration() {

	const now = 1589570165

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	s = strings.ReplaceAll(s, " мин. ", "m")
	s = strings.ReplaceAll(s, " сек.", "s")

	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	nano := int64(d.Seconds())

	fmt.Println(time.Unix(now+nano, 0).UTC().Format(time.UnixDate))
}

func DurationBetweenDate() {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	sl := strings.Split(strings.TrimSuffix(s, "\n"), ",")

	t1, err := time.Parse("02.01.2006 15:04:05", sl[0])
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("02.01.2006 15:04:05", sl[1])
	if err != nil {
		panic(err)
	}

	r := t1.Sub(t2)
	if r < 0 {
		r = t2.Sub(t1)
	}

	dur, err := time.ParseDuration(r.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(dur.String())
}

func ReadJsonFromFile() {
	file, err := os.Open("")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	var props []Properties

	body, err := io.ReadAll(file)

	err = json.Unmarshal(body, &props)
	if err != nil {
		fmt.Println(err.Error())
	}
	sum := int64(0)
	for _, v := range props {
		sum += int64(v.GlobalID)
	}
	fmt.Println(sum)
}

type Properties struct {
	GlobalID int `json:"global_id"`
}

func AveregeStudentsRating() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	var (
		group Group
		aver  float64
	)
	err = json.Unmarshal(in, &group)
	if err != nil {
		return
	}

	for _, student := range group.Students {
		aver += float64(len(student.Rating))
	}

	out := struct {
		Average float64
	}{
		Average: roundFloat(aver/float64(len(group.Students)), 1),
	}

	res, _ := json.MarshalIndent(out, "", "    ")

	fmt.Printf("%s", res)
}

type Group struct {
	ID       int        `json:"ID"`
	Number   string     `json:"Number"`
	Year     int        `json:"Year"`
	Students []Students `json:"Students"`
}
type Students struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ReadWrite() {
	var (
		s   string
		sum = 0
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = scanner.Text()
		if v, err := strconv.Atoi(s); err == nil {
			sum += v
			continue
		}
		break
	}

	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(strconv.Itoa(sum))
	if err != nil {
		return
	}
	writer.Flush()
}

type Battery struct {
	Power string
}

func NewBattery(power string) *Battery {
	return &Battery{Power: power}
}

func (b Battery) String() string {
	zero := strings.Count(b.Power, "0")
	return fmt.Sprintf("[%s%s]", strings.Repeat(" ", zero), strings.Repeat("X", 10-zero))
}

func BatteryRunnner() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	batteryForTest := NewBattery(str)
	fmt.Println(batteryForTest.String())
}

func Calculator() {

	value1, value2, operation := readTask()

	operands := []float64{}
	var result float64

	for _, v := range []interface{}{value1, value2} {
		switch val := v.(type) {
		case float64:
			operands = append(operands, val)
		default:
			fmt.Printf("value=%v:%T\n", val, val)
			return
		}
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result = operands[0] + operands[1]
		case "-":
			result = operands[0] - operands[1]
		case "*":
			result = operands[0] * operands[1]
		case "/":
			result = operands[0] / operands[1]
		default:
			fmt.Println("неизвестная операция")
		}
		fmt.Printf("%.4f", result)
	}
}

func readTask() (interface{}, interface{}, interface{}) {
	return 2, 1, "+"
}

func RemoveOdds() {
	fn := func(n uint) uint {
		s := fmt.Sprint(n)
		r := ""
		for _, b := range []byte(s) {
			fmt.Println(b - '0')
			if (b-'0')%2 == 0 {
				r += string(b)
			}
		}
		res, _ := strconv.ParseUint(r, 10, 0)
		if res == 0 {
			return 100
		}
		return uint(res)
	}
	fmt.Println(fn(727178))
}

func work(x int) int {
	return x + 1
}

func adding(s1, s2 string) int64 {
	var n1, n2 int64
	s3 := ""
	for _, v := range []rune(s1) {
		if unicode.IsDigit(v) {
			s3 += string(v)
		}
	}
	n1, _ = strconv.ParseInt(s3, 10, 0)
	s3 = ""
	for _, v := range []rune(s2) {
		if unicode.IsDigit(v) {
			s3 += string(v)
		}
	}
	n2, _ = strconv.ParseInt(s3, 10, 0)
	return n1 + n2
}
