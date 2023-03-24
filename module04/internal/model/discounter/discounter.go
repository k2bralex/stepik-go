package discounter

type Discounter interface {
	CalcDiscount(d int) (float64, error)
}
