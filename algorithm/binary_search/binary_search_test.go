package binary_search

import "testing"

var (
	slice  = []int{1, 3, 5, 7, 77}
	target = 5
	expect = 2
)

func Test_BinarySearch(t *testing.T) {
	index := binarySearch(slice, target, 0, len(slice)-1)
	if index != expect {
		t.Fatal("error")
	}
	t.Log(index)
}

func Test_BinarySearch2(t *testing.T) {
	index := binarySearch2(slice, target, 0, len(slice)-1)
	if index != expect {
		t.Fatal("error")
	}
	t.Log(index)
}

func Benchmark_BinarySearch(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		binarySearch(slice, target, 0, len(slice)-1)
	}
}

func Benchmark_BinarySearch2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		binarySearch2(slice, target, 0, len(slice)-1)
	}
}
