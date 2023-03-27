package main

import (
	"fmt"
	"udmey/module06/internal/util"
)

func main() {
	r, err := util.ReverseNumber(12)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
}
