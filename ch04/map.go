package main

import "fmt"

func main() {
	//ages := make(map[string]int)

	ages := map[string]int{
		"alice":   22,
		"charlie": 22,
	}

	fmt.Println(ages["alice"])

	delete(ages, "alice")

	fmt.Println(ages["alice"])

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var ages1 map[string]int
	fmt.Println(ages1 == nil)
	fmt.Println(len(ages1) == 0)

	ages1["carol"] = 21
}
