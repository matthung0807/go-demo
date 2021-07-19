package serivce

type CalculatorService interface {
	Plus(int, int) int
}

type Calculator struct {
}

func (c Calculator) Plus(x, y int) int {
	return x + y
}
