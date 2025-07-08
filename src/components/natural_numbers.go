package components

type NaturalNumber struct {
	Value uint64
}

func (n NaturalNumber) IsZero() bool {
	return n.Value == 0
}

func (n NaturalNumber) Add(other NaturalNumber) NaturalNumber {
	if other.IsZero() {
		return n
	}

	return S(n.Add(P(other)))
}

func Int(value uint64) NaturalNumber {
	return NaturalNumber{
		Value: value,
	}
}
