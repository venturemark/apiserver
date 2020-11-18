package data

type wrapper struct {
	space string
	value []int64
}

func (w wrapper) GetSpace() string {
	return w.space
}

func (w wrapper) GetValue() []int64 {
	return w.value
}
