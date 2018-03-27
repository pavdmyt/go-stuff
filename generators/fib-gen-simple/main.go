package main

import "fmt"

func fibGen() <-chan int {
	out := make(chan int)

	go func() {
		a, b := 0, 1
		for {
			out <- a
			a, b = b, a+b
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range fibGen() {
		fmt.Println(n)
		if n > 100000 {
			break
		}
	}
}
