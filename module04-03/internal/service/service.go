package service

import (
	"errors"
	"fmt"
	"udmey/module04-03/internal/model"
)

const DEFAULT_DISCOUNT = 20

func Run() {
	cust := model.NewCustomer("Alex", 20, 20000.32, 0, false)

	price, err := CalcPrice(cust, 399.99)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Final price: %.2f\n", price)
}

func CalcPrice(obj model.Discounter, price float64) (float64, error) {
	discount, err := obj.CalcDiscount(DEFAULT_DISCOUNT)
	if err != nil {
		return price, err
	}
	if discount > price {
		return price, errors.New("discount more than price")
	}
	return price - discount, nil
}
