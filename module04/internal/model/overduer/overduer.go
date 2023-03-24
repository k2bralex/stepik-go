package overduer

type Overduer struct {
	Balance float64
	Debt    float64
}

func NewOverduer(balance, debt float64) *Overduer {
	return &Overduer{
		Balance: balance,
		Debt:    debt,
	}
}
