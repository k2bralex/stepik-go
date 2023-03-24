package partner

import "errors"

type Partner struct {
	Name    string
	Age     int
	balance float64
	debt    float64
}

func NewPartner(name string, age int, balance, debt float64) *Partner {
	return &Partner{
		Name:    name,
		Age:     age,
		balance: balance,
		debt:    debt,
	}
}

func (p *Partner) WrOffDebt() error {
	if p.balance > 50000.00 {
		p.debt = 0
		return nil
	}
	return errors.New("low balance")
}
