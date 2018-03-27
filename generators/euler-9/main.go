package main

import (
	"fmt"
	"math"
)

// Pair stores factor-pairs
type Pair struct {
	a int
	b int
}

// Triple stores Pythagorean triples
type Triple struct {
	a int
	b int
	c int
}

// Generates factor-pairs of the given number.
// E.g. factor-pairs of 18 are: (1, 18), (2, 9), (3, 6)
func genDivisorPairs(num int) <-chan Pair {
	out := make(chan Pair)

	go func() {
		numSqrt := math.Sqrt(float64(num))
		for divisor := 1; divisor < int(numSqrt+1); divisor++ {
			if num%divisor == 0 {
				out <- Pair{divisor, num / divisor}
			}
		}
		close(out)
	}()
	return out
}

// Generates Pythagorean triples using Dickson's method.
//
// To find triples: a^2 + b^2 = c^2,
// find ints r, s and t such that: r^2 = 2*s*t
// Then:
//   a = r + s, b = r + t, c = r + s + t
func genTriples() <-chan Triple {
	out := make(chan Triple)

	go func() {
		r := 2 // any even int
		for {
			st := (r * r) / 2
			for pair := range genDivisorPairs(st) {
				s, t := pair.a, pair.b
				triple := Triple{r + s, r + t, r + s + t}
				out <- triple
			}
			r += 2
		}
		close(out)
	}()
	return out
}

func main() {
	for triple := range genTriples() {
		a, b, c := triple.a, triple.b, triple.c
		if a+b+c == 1000 {
			fmt.Println(a * b * c)
			break
		}
	}
}
