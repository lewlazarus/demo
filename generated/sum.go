// Generated code

package generated

func Sum(values []float64) float64 {
	var res float64 = 0
	var c int

	l := len(values)
	i := 0

	c = (l - i) / 10
	if c > 0 {
		for j := 0; j < c; j++ {
			res += values[i] + values[i+1] + values[i+2] + values[i+3] + values[i+4] + values[i+5] + values[i+6] + values[i+7] + values[i+8] + values[i+9]
			i += 10
		}
	}

	c = (l - i) / 1
	if c > 0 {
		for j := 0; j < c; j++ {
			res += values[i]
			i += 1
		}
	}

	return res
}
