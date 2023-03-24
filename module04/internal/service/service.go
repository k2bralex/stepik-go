package service

import (
	"errors"
	"fmt"
	"udmey/module04/internal/model/customer"
	"udmey/module04/internal/model/debtor"
	"udmey/module04/internal/model/discounter"
	"udmey/module04/internal/model/partner"
)

const DEFAULT_DISCOUNT = 20

func Run() {
	cust := customer.NewCustomer("Alex", 20, 20000.32, 0, false)
	part := partner.NewPartner("John", 40, 78000, 300)

	fmt.Println(cust)
	fmt.Println(part)

	if err := startDynamicTransaction(cust); err != nil {
		fmt.Println(err.Error())
	}

	if err := startDynamicTransaction(part); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cust)
	fmt.Println(part)
}

func startDynamicTransaction(t interface{}) error {
	switch t.(type) {
	case *partner.Partner:
		return t.(debtor.Debtor).WrOffDebt()
	default:
		return errors.New("incorrect type")
	}
}

func calcPrice(obj discounter.Discounter, price float64) (float64, error) {
	discount, err := obj.CalcDiscount(DEFAULT_DISCOUNT)
	if err != nil {
		return price, err
	}
	if discount > price {
		return price, errors.New("discount more than price")
	}
	return price - discount, nil
}
