package main

import (
	"fmt"
	"udmey/module06/internal"
)

func main() {
	r, err := internal.ReverseNumber(12)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
}
