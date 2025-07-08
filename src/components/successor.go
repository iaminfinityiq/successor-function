package components

type SuccessorTable struct {
	Map      map[NaturalNumber]NaturalNumber
	Inverted map[NaturalNumber]NaturalNumber
}

func (s SuccessorTable) access_successor(n NaturalNumber) NaturalNumber {
	return s.Map[n]
}

func (s SuccessorTable) access_inverted(n NaturalNumber) NaturalNumber {
	return s.Inverted[n]
}

func Successor() SuccessorTable {
	var returned SuccessorTable = SuccessorTable{
		Map:      make(map[NaturalNumber]NaturalNumber),
		Inverted: make(map[NaturalNumber]NaturalNumber),
	}

	var prev uint64 = 0
	for i := uint64(1); i < 1000000; i++ {
		returned.Map[Int(prev)] = Int(i)
		returned.Inverted[Int(i)] = Int(prev)
		prev = i
	}

	return returned
}

var successor_table SuccessorTable = Successor()

func S(n NaturalNumber) NaturalNumber {
	return successor_table.access_successor(n)
}

func SInverted(n NaturalNumber) NaturalNumber {
	return successor_table.access_inverted(n)
}
