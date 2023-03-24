package customer

import (
	"errors"
	. "udmey/module04/internal/model/overduer"
)

type Customer struct {
	Name     string
	Age      int
	Discount bool
	Overduer *Overduer
}

func NewCustomer(name string, age int, balance, debt float64, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Discount: discount,
		Overduer: NewOverduer(balance, debt),
	}
}

func (c *Customer) CalcDiscount(disc int) (float64, error) {
	if !c.Discount {
		return 0, errors.New("discount not available")
	}
	result := float64(disc) - c.Overduer.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (c *Customer) WrOffDebt() error {
	if c.Overduer.Balance > 70000.00 {
		c.Overduer.Debt = 0
		return nil
	}
	return errors.New("low balance")
}
