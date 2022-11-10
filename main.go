package main

import (
	"fmt"
	"github.com/asistent0/sf_basics_go/maths"
)

func main() {
	print(maths.Sum, 2, 3)
	print(maths.Sub, 5, 3)
}

func print(f func(int, int) int, a int, b int) {
	fmt.Println(f(a, b))
}
