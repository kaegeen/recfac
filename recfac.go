package main

import "fmt"

func main() {

	fmt.Println(Factorial(4))

}

func Factorial(n int) int {

	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return 1
	}

	result := Factorial(n - 1)
	if result > 0 && n > (1<<31-1)/result {
		return 0
	}
	return n * result
}
