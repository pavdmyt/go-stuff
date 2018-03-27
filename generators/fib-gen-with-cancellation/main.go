package main

/*
Discussion on generators and implementing cancellation:
https://stackoverflow.com/a/34466755
*/

import "fmt"

func fibGen(abort <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		a, b := 0, 1
		for {
			select {
			case out <- a:
				a, b = b, a+b
			case <-abort:
				return
			}
		}
	}()
	return out
}

func main() {
	abort := make(chan struct{})

	for n := range fibGen(abort) {
		fmt.Println(n)
		if n > 100000 {
			close(abort)
			break
		}
	}

}
