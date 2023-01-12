package main

import "fmt"

func main() {
	fmt.Println(throwEggs(4, 130))
}

func throwEggs(eggs, floors uint) uint {
	var (
		windows        = [2][]uint{make([]uint, eggs), make([]uint, eggs)}
		testNum   uint = 0
		curFloors uint = 0
	)

	for curFloors+1 < floors {
		testNum++
		windows[1][0] = testNum
		for i := 1; i < len(windows[1]); i++ {
			windows[1][i] = windows[0][i-1] + windows[0][i] + 1
		}
		curFloors = windows[1][eggs-1]
		windows[0], windows[1] = windows[1], windows[0]
	}
	fmt.Println(windows[0][eggs-1], " ", testNum)

	return testNum
}
