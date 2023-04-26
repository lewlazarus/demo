package ops

import "testing"

func BenchmarkMultiply1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(v1000)
	}
}
