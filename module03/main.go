package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("/Users/oleksiibrahanets/GolandProjects/udemy/module03/data.json")
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
