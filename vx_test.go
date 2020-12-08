package vx

import (
	"testing"
)

func TestMalloc(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			defer Free(x)

			if len(x) != align(size) {
				t.Errorf("Malloc should return float slice size of %d, but size is %d", align(size), len(x))
			}
		}(size)
	}
}

func TestAdd(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			y := Malloc(size)
			z := Malloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] + y[i]
			}

			Add(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Add should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestSub(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			y := Malloc(size)
			z := Malloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] - y[i]
			}

			Sub(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Mul should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestMul(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			y := Malloc(size)
			z := Malloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] * y[i]
			}

			Mul(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Mul should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestDiv(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			y := Malloc(size)
			z := Malloc(size)
			defer Free(x)
			defer Free(y)
			defer Free(z)

			truth := make([]float32, size)
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth[i] = x[i] / y[i]
			}

			Div(size, x, y, z)

			for i := 0; i < size; i++ {
				if truth[i] != z[i] {
					t.Errorf("Mul should return %f in %d, but %f", truth[i], i, z[i])
				}
			}
		}(size)
	}
}

func TestDot(t *testing.T) {
	for _, size := range []int{7, 8, 15} {
		func(size int) {
			x := Malloc(size)
			y := Malloc(size)
			defer Free(x)
			defer Free(y)

			var truth float32
			for i := 0; i < size; i++ {
				x[i] = float32(i)
				y[i] = float32(i + 1)
				truth += x[i] * y[i]
			}

			result := Dot(size, x, y)
			if truth != result {
				t.Errorf("Dot should return %f, but %f", truth, result)
			}
		}(size)
	}
}

func TestAlign(t *testing.T) {
	expects := [][]int{
		{7, 8},
		{8, 8},
		{9, 16},
	}
	for _, expect := range expects {
		if size := align(expect[0]); size != expect[1] {
			t.Errorf("align should return %d, but %d", expect[1], size)
		}
	}
}