package calc

import "math"

func Variance(mean float64, values []float64) float64 {
	var res float64 = 0
	var c int

	l := len(values)
	i := 0

	c = (l - i) / 10
	if c > 0 {
		for j := 0; j < c; j++ {
			res += math.Pow(values[i]-mean, 2) + math.Pow(values[i+1]-mean, 2) + math.Pow(values[i+2]-mean, 2) + math.Pow(values[i+3]-mean, 2) + math.Pow(values[i+4]-mean, 2) + math.Pow(values[i+5]-mean, 2) + math.Pow(values[i+6]-mean, 2) + math.Pow(values[i+7]-mean, 2) + math.Pow(values[i+8]-mean, 2) + math.Pow(values[i+9]-mean, 2)
			i += 10
		}
	}

	c = (l - i) / 1
	if c > 0 {
		for j := 0; j < c; j++ {
			res += math.Pow(values[i]-mean, 2)
			i += 1
		}
	}

	return res
}
