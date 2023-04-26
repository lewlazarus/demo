package ops

import "testing"

func BenchmarkMean1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(v1000)
	}
}
