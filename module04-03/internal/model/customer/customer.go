package customer

import "errors"

type Customer struct {
	Name     string
	Age      int
	Balance  float64
	Debt     float64
	Discount bool
}

func NewCustomer(name string, age int, balance, debt float64, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func (c *Customer) CalcDiscount(disc int) (float64, error) {
	if !c.Discount {
		return 0, errors.New("discount not available")
	}
	result := float64(disc) - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}
