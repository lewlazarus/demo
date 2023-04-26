// Generated code

package generated

import "math"

func Variance(values []float64) float64 {
	l := len(values)
	if l <= 1 {
		return 0.0
	}

	var res float64 = 0
	var c int

	m := Mean(values)
	i := 0

	c = (l - i) / 10
	if c > 0 {
		for j := 0; j < c; j++ {
			res += math.Pow(values[i]-m, 2) + math.Pow(values[i+1]-m, 2) + math.Pow(values[i+2]-m, 2) + math.Pow(values[i+3]-m, 2) + math.Pow(values[i+4]-m, 2) + math.Pow(values[i+5]-m, 2) + math.Pow(values[i+6]-m, 2) + math.Pow(values[i+7]-m, 2) + math.Pow(values[i+8]-m, 2) + math.Pow(values[i+9]-m, 2)
			i += 10
		}
	}

	c = (l - i) / 1
	if c > 0 {
		for j := 0; j < c; j++ {
			res += math.Pow(values[i]-m, 2)
			i += 1
		}
	}

	return res / float64(l)
}
