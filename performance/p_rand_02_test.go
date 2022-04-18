package performance

import (
	"fmt"
	"testing"
)

func Test_rand01(t *testing.T) {
	fmt.Println(rand01(32))
}
func Test_rand02(t *testing.T) {
	fmt.Println(rand02(32))
}
func BenchmarkRand01(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand01(32)
		}
	})
}
func BenchmarkRand02(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand02(32)
		}
	})
}
