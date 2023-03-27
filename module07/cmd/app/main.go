package main

import (
	"fmt"
	"udmey/module07/internal/service"
)

func main() {
	v := service.StructToMap(struct {
		name string
		age  int
	}{name: "Alex", age: 22})
	fmt.Println(v)
}
