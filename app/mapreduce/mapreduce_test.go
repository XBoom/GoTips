package mapreduce

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {

		Run(uids []int)(30) // run fib(30) b.N times
	}
}