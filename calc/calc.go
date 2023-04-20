package calc

type Calculater interface {
	Addition() int
	AdditionFloat() int
}

type MyCalculater struct{}

type Input struct {
	Num1 int
	Num2 int
}

type Inputfloat struct {
	Numfloat1 float64
	Numfloat2 float64
}

func (data *Input) Addition() int {
	return data.Num1 + data.Num2

}
func (datafloat *Inputfloat) AdditionFloat() float64 {
	return datafloat.Numfloat1 + datafloat.Numfloat2

}

func (data *Input) Substraction() int {
	return data.Num1 - data.Num2

}

func (data *Input) Multiplication() int {
	return data.Num1 * data.Num2

}

func (data *Input) Division() int {
	return data.Num1 / data.Num2

}
