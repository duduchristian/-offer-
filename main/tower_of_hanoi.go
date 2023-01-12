package main

import "fmt"

func main() {
	towerOfHanoi(4, "A", "B", "C")
}

func towerOfHanoi(num int, start, middle, to string) {
	if num > 0 {
		towerOfHanoi(num-1, start, to, middle)
		fmt.Printf("move %s to %s.\n", start, to)
		towerOfHanoi(num-1, middle, start, to)
	}
}

func towerOfHanoiStack(num int, start, middle, to string) {
	//stack := make([]string, 0,
}
