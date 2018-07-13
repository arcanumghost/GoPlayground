package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 5
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func fizzbuzz(i int) {
	if i%15 == 0 {
		fmt.Println("fizzbuzz")
	} else if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(i)
	}
}
func main() {
	for i := 1; i <= 100; i++ {
		fizzbuzz(i)
	}
}
