package ops

import "testing"

func BenchmarkSum1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(v1000)
	}
}
