package main

import (
	"fmt"

	"./typedata"
)

func main() {

	a, b := swap(1, 2)

	fmt.Printf("%d\n", a)

	fmt.Printf("%d\n", b)

	if c, d := swap(3, 4); c > 3 {
		fmt.Println(c)
	} else {
		fmt.Println(2 * d)
	}

	typedata.Typedata()
}

func swap(x, y int) (int, int) {
	return y, x
}
