package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Welcome to begin")
}

func main() {
	x := 10
	y := 15
	fmt.Println("Hello")
	fmt.Println(x + y)

	if x > y {
		fmt.Println("You are correct")
	} else {
		fmt.Println("exit")
	}
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and type %T", z, x)
	} else {
		fmt.Printf("z is %v and not type %T\n", z, x)
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	d1 := time.Duration(rand.Int63n(230))
	d2 := time.Duration(rand.Int31n(250))
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42

	}()
	select {
	case v1 := <-ch1:
		fmt.Println("value from the channel 1:", v1)
	case v2 := <-ch2:
		fmt.Println("value from the channel 2:", v2)
	}
}
