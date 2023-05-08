package process

import (
	"demo/process/calc"
	"fmt"
	"testing"
)

var v10 = []float64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

var v100 = []float64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
}

func TestOps(t *testing.T) {
	testCases := []struct {
		name     string
		opFn     func([]float64) float64
		values   []float64
		expected float64
	}{
		{"Sum", opSum, v10, 55},
		{"Multiply", opMultiply, v10, 3628800.0},
		{"Mean", opMean, v10, 5.5},
		{"Variance", opVariance, v10, 9.1666667},
		{"StandardDeviation", opStandardDeviation, v10, 3.027650},
	}

	for _, tc := range testCases {
		expected := fmt.Sprintf("%.5f", tc.expected)
		res := fmt.Sprintf("%.5f", tc.opFn(tc.values))

		if res != expected {
			t.Errorf("%s: expected %s, result %s", tc.name, expected, res)
		}
	}
}

func TestProcess(t *testing.T) {
	testCases := []struct {
		values   []float64
		expected float64
	}{
		{v10, 55},
		{v100, 5050},
	}

	for _, tc := range testCases {
		res := process(calc.Sum, sum, 0, tc.values)

		if res != tc.expected {
			t.Errorf("expected %f, result %f", tc.expected, res)
		}
	}
}

func BenchmarkSum100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opSum(v100)
	}
}

func BenchmarkMultiply100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opMultiply(v100)
	}
}

func BenchmarkMean100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opMean(v100)
	}
}

func BenchmarkVariance100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opVariance(v100)
	}
}

func BenchmarkStandardDeviation100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		opStandardDeviation(v100)
	}
}
