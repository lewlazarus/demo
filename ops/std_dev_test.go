package ops

import "testing"

func BenchmarkStdDev1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardDeviation(v1000)
	}
}
