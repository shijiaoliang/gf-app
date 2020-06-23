package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 7, 77}
	target := 5
	expect := 2

	index := binarySearch(a, target, 0, len(a)-1)
	if index != expect {
		t.Fatal("error")
	}
	t.Log(index)
}

func TestBinarySearch2(t *testing.T) {
	a := []int{1, 3, 5, 7, 77}
	target := 5
	expect := 2

	index := binarySearch2(a, target, 0, len(a)-1)
	if index != expect {
		t.Fatal("error")
	}
	t.Log(index)
}

func BenchmarkBinarySearch(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	a := []int{1, 3, 5, 7, 77}
	target := 5
	for i := 0; i < b.N; i++ {
		binarySearch(a, target, 0, len(a)-1)
	}
}

func BenchmarkBinarySearch2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	a := []int{1, 3, 5, 7, 77}
	target := 5
	for i := 0; i < b.N; i++ {
		binarySearch2(a, target, 0, len(a)-1)
	}
}
