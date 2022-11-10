package main

import (
	"fmt"
	"github.com/asistent0/sf_basics_go/maths"
)

func main() {
	myPrint(maths.Sum, 2, 3)
	myPrint(maths.Sub, 5, 3)
	myPrint(func(a int, b int) int { return a * b }, 5, 3)
}

func myPrint(f func(int, int) int, a int, b int) {
	fmt.Println(f(a, b))
}
