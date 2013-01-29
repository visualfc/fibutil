fibutil
=======

Golang Fibonacci numbers library

Fibonacci numbers integer sequence:

	0 1 1 2 3 5 8 13 21 34 55 89 144 ...
	-21 13 -8 5 －3 2 －1 1 0 1 1 2 3 5 8 13 21 ...

### Install

	go get -v github.com/visualfc/fibutil

### fibutil/fib

	func fib.Fib(n int64) *big.Int
	func fib.FastFib(n int64) *big.Int
	func fib.FibList(n1, n2 int64) []string
	
### fibutil	
Fibonacci number util

Usage:

	fibutil n		:fibonacci number
	fibutil n1 n2	:fibonacci number list

Example:	

	fibutil 10 		=> 55
	fibutil -10 10  => [-55 34 -21 13 -8 5 -3 2 -1 1 0 1 1 2 3 5 8 13 21 34 55]
	
