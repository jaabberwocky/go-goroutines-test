package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {

	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func printstuff(stuff string) {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Millisecond)
	fmt.Println(stuff)
}

func main() {
	f("blocking")

	go f("first")
	go f("second")
	go f("third")
	go printstuff("stuff")

	fmt.Scanln()
	fmt.Println("done")
}
