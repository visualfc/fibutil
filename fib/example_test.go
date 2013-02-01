// Copyright 2013 visualfc <visualfc@gmail.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fib

import (
	"fmt"
)

func ExampleFibList() {
	fmt.Println(FibList(-10, 10))
	// Output: [-55 34 -21 13 -8 5 -3 2 -1 1 0 1 1 2 3 5 8 13 21 34 55]
}
