package mapreduce

import (
	"flag"
	"os"
	"testing"
)

var data []int
const sliceSize = 10000

func TestMain(m *testing.M) {
	flag.Parse()
	data = make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		data[i] = i
	}
	os.Exit(0)
}

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run(data)// run fib(30) b.N times
	}
}

func BenchmarkRun1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run1(data)// run fib(30) b.N times
	}
}