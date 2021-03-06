package comparison

type Result uint8

const (
	FirstBigger Result = 1 << iota
	Equal
	SecondBigger

	FirstSmaller  = SecondBigger
	SecondSmaller = FirstBigger
)

func FromInt(result int) Result {
	switch {
	case result > 0:
		return SecondBigger
	case result < 0:
		return FirstBigger
	default:
		return Equal
	}
}
