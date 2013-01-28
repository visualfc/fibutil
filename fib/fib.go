// fibonacci number
//
// 0 1 1 2 3 5 8 13 21 34 55 89 144 ...
//
// -21 13 -8 5 3 2 1 1 0 1 1 2 3 5 8 13 21 ...
package fib

import (
	"fmt"
	"math/big"
)

// 非递归方式实现Fibonacci算法
func Fibm(n int64) int64 {
	if n == 0 {
		return 0
	}
	var i, a, b int64
	b = 1
	if n < 0 {
		n = -n
		if n%2 == 0 {
			b = -1
		}
	}
	for i = 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 递归方式实现Fibonacci算法
func Fibm2(n int64) int64 {
	if n < 0 {
		if n%2 == 0 {
			return _Fibm2(-n, 0, -1)
		} else {
			return _Fibm2(-n, 0, 1)
		}
	}
	return _Fibm2(n, 0, 1)
}

func _Fibm2(n, a1, a2 int64) int64 {
	if n == 0 {
		return a1
	}
	return _Fibm2(n-1, a2, a1+a2)
}

// big.Int实现Fibonacci算法
func Fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	a, b, c := big.NewInt(0), big.NewInt(1), big.NewInt(0)
	if n < 0 {
		n = -n
		if n%2 == 0 {
			b.SetInt64(-1)
		}
	}
	for i := int64(2); i <= n; i++ {
		c.Set(a)
		a.Set(b)
		b.Add(b, c)
	}
	return b
}

type Matrix struct {
	M00, M01, M10, M11 *big.Int
}

func NewMatrix(M00, M01, M10, M11 int64) *Matrix {
	return &Matrix{big.NewInt(M00), big.NewInt(M01), big.NewInt(M10), big.NewInt(M11)}
}

func (m *Matrix) String() string {
	return fmt.Sprintf("[%s,%s;%s;%s]", m.M00, m.M01, m.M10, m.M11)
}

func (m *Matrix) Mul(other *Matrix) *Matrix {
	var a, a1, b, b1, c, c1, d, d1 big.Int
	return &Matrix{
		a.Mul(m.M00, other.M00).Add(&a, a1.Mul(m.M01, other.M10)),
		b.Mul(m.M00, other.M01).Add(&b, b1.Mul(m.M01, other.M11)),
		c.Mul(m.M10, other.M00).Add(&c, c1.Mul(m.M11, other.M10)),
		d.Mul(m.M10, other.M01).Add(&d, d1.Mul(m.M11, other.M11))}
}

func (m *Matrix) Pow(n int64) *Matrix {
	if n == 1 {
		return m
	} else if n%2 == 0 {
		out := m.Pow(n / 2)
		return out.Mul(out)
	} else {
		out := m.Pow((n - 1) / 2)
		out = out.Mul(out)
		return out.Mul(m)
	}
	return nil
}

// 使用矩阵求解Fibonacci算法
func FastFib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 || n == -1 {
		return big.NewInt(1)
	}
	m := NewMatrix(1, 1, 1, 0)
	if n < 0 {
		n = -n
		if n%2 == 0 {
			m.M00.SetInt64(-1)
		}
	}
	return m.Pow(n - 1).M00
}

// big.Int实现Fibonacci列表
func FibList(n1, n2 int64) []string {
	size := n2 - n1
	if size < 0 {
		return nil
	} else if size == 0 {
		return []string{FastFib(n1).String()}
	}
	a := FastFib(n1)
	b := FastFib(n1 + 1)
	r := make([]string, size+1)
	r[0] = a.String()
	r[1] = b.String()
	var tmp big.Int
	for i := int64(2); i <= size; i++ {
		tmp.Set(a)
		a.Set(b)
		b.Add(b, &tmp)
		r[i] = b.String()
	}
	return r
}
