// fibutil project main.go
package main

import (
	"fmt"
	"github.com/visualfc/fibutil/fib"
	"os"
	"strconv"
)

func main() {
	switch len(os.Args) {
	case 2:
		n, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			fmt.Println(fib.FastFib(n))
			return
		}
	case 3:
		n1, e1 := strconv.ParseInt(os.Args[1], 10, 64)
		n2, e2 := strconv.ParseInt(os.Args[2], 10, 64)
		if e1 == nil && e2 == nil {
			fmt.Println(fib.FibList(n1, n2))
			return
		}
	}
	fmt.Fprintf(os.Stderr, "%s, Fibonacci number util\n\tfibutil n\t: Fibonacci number\n\tfibutil n1 n2\t: Fibonacci number list\n", os.Args[0])
}
