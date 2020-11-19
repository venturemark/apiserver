package data

type Wrapper struct {
	Space string
	Value []float64
}

func (w Wrapper) GetSpace() string {
	return w.Space
}

func (w Wrapper) GetValue() []float64 {
	return w.Value
}
