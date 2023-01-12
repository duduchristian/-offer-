package main

import "fmt"

func main() {
	base := 100000.0
	month := 21 * 12
	for i := 0; i < month; i++ {
		base += base * 0.03 / 12
	}

	res1 := 2879.0*329 - base + 100000
	res2 := 2315.0*329 + 100000
	fmt.Println(res1, " ", res2)
}
