package fib

import (
	"testing"
)

func TestFibm(t *testing.T) {
	ar := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i < len(ar); i++ {
		if ar[i] != Fibm(int64(i)) {
			t.Fatalf("%d, %d != %d", i, ar[i], Fibm(int64(i)))
		}
	}
	ar1 := []int64{0, 1, -1, 2, -3, 5, -8, 13, -21}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != Fibm(int64(-i)) {
			t.Fatalf("%d, %d != %d", -i, ar1[i], Fibm(int64(-i)))
		}
	}
}

func TestFibm2(t *testing.T) {
	ar := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i < len(ar); i++ {
		if ar[i] != Fibm2(int64(i)) {
			t.Fatalf("%d, %d != %d", i, ar[i], Fibm2(int64(i)))
		}
	}
	ar1 := []int64{0, 1, -1, 2, -3, 5, -8, 13, -21}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != Fibm2(int64(-i)) {
			t.Fatalf("%d, %d != %d", -i, ar1[i], Fibm2(int64(-i)))
		}
	}
}

func TestFib(t *testing.T) {
	ar := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i < len(ar); i++ {
		if ar[i] != Fib(int64(i)).Int64() {
			t.Fatalf("%d, %d != %d", i, ar[i], Fib(int64(i)).Int64())
		}
	}
	ar1 := []int64{0, 1, -1, 2, -3, 5, -8, 13, -21}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != Fib(int64(-i)).Int64() {
			t.Fatalf("%d, %d != %d", -i, ar1[i], Fib(int64(-i)))
		}
	}
}

func TestFastFib(t *testing.T) {
	ar := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i < len(ar); i++ {
		if ar[i] != FastFib(int64(i)).Int64() {
			t.Fatalf("%d, %d != %d", i, ar[i], FastFib(int64(i)))
		}
	}
	ar1 := []int64{0, 1, -1, 2, -3, 5, -8, 13, -21}
	for i := 0; i < len(ar1); i++ {
		if ar1[i] != FastFib(int64(-i)).Int64() {
			t.Fatalf("%d, %d != %d", -i, ar1[i], FastFib(int64(-i)))
		}
	}
}

func BenchmarkFibm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibm(90)
	}
}

func BenchmarkFibm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibm2(90)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(90)
	}
}

func BenchmarkFastFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFib(90)
	}
}

func BenchmarkFib200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(200)
	}
}

func BenchmarkFastFib200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFib(200)
	}
}

func BenchmarkFib1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(1000)
	}
}

func BenchmarkFastFib1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFib(1000)
	}
}
