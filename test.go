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

	ns := []int{10, 20, 30, 40, 50} // 要素にアクセス
	fmt.Println(cap(ns))
	fmt.Println(ns[3])
	// 長さ
	fmt.Println(len(ns))
	// 要素􏰁追加
	// 容量が足りない場合􏰀背後􏰁配列が再確保される
	ns = append(ns, 60, 70)
	// 容量
	fmt.Println(cap(ns))

	ns = append(ns, 60, 70)

	fmt.Println(cap(ns))

	ns = append(ns, 60, 70)

	fmt.Println(cap(ns))
}

func swap(x, y int) (int, int) {
	return y, x
}
