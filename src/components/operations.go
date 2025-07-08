package components

func Plus(a NaturalNumber, b NaturalNumber) NaturalNumber {
	if b.IsZero() {
		return a
	}

	return S(Plus(a, P(b)))
}

func Monus(a NaturalNumber, b NaturalNumber) NaturalNumber {
	if b.IsZero() {
		return a
	}

	return Monus(P(a), P(b))
}
