package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Shownum() {
	for i :=0; i<100; i++ {
		foo := rand.Intn(100)
		fmt.Println(foo)
		if foo < 10 {
			break
		}
	}
}

func main() {
	now := time.Now()
	secs := now.Unix()
	rand.Seed(secs)
	fmt.Println("dude!")
	Shownum()
}

