package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("first")
	go f("second")
	
	fmt.Scanln()
	fmt.Println("done")
}
