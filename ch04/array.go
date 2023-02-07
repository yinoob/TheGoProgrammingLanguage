package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Println("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Println("%d\n", v)
	}

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	a1 := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a1 == b, a1 == c, b == c) // "true false false"

}
