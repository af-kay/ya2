package ya2

type figure int

const (
	square figure = iota
	circle
	triangle
)

func area(fig figure) (fn func(float64) float64, ok bool) {
	switch fig {
	case square:
		fn = func(side float64) float64 { return side * side }
	case circle:
		fn = func(radius float64) float64 { return 3.14 * radius * radius }
	case triangle:
		fn = func(side float64) float64 { return (side * 3) / 2 }
	}

    return fn, fn != nil
}
