// Generated code

package generated

func Mean(values []float64) float64 {
	l := len(values)

	switch l {
	case 0:
		return 0
	case 1:
		return values[0]
	}

	return Sum(values) / float64(l)
}
