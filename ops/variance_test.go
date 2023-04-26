package ops

import "testing"

func BenchmarkVariance1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Variance(v1000)
	}
}
