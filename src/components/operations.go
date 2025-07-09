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

func Times(a NaturalNumber, b NaturalNumber) NaturalNumber {
	if b.IsZero() {
		return b
	}

	return Plus(a, Times(a, P(b)))
}

func Divide(a NaturalNumber, b NaturalNumber) NaturalNumber {
	if b.IsZero() {
		return b
	}

	var monus NaturalNumber = Monus(S(a), b)
	if monus.IsZero() {
		return zero
	}

	return S(Divide(monus, b))
}

func Mod(a NaturalNumber, b NaturalNumber) NaturalNumber {
	return Monus(a, Times(Divide(a, b), b))
}

func Exponents(a NaturalNumber, b NaturalNumber) NaturalNumber {
	if b.IsZero() {
		return one
	}

	return Times(a, Exponents(a, P(b)))
}
