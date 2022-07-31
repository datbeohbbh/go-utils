package interfaces

type signedInteger interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type Integer[T signedInteger] struct {
	Value T
}

func (t Integer[signedInteger]) Less(rhs any) bool {
	x := rhs.(Integer[signedInteger])
	return (t.Value < x.Value)
}

func (t Integer[signedInteger]) Equal(rhs any) bool {
	x := rhs.(Integer[signedInteger])
	return (t.Value == x.Value)
}
