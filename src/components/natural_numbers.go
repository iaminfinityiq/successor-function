package components

type NaturalNumber struct {
	Value uint64
}

func (n NaturalNumber) IsZero() bool {
	return n == zero
}

func Int(value uint64) NaturalNumber {
	return NaturalNumber{
		Value: value,
	}
}
