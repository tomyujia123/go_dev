package main

import "fmt"

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func concat(a string, arg ...string) (result string) {
	result = a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return result
}

func main() {
	sum := add(10, 3, 5)
	fmt.Println(sum)

	result := concat("hello", " ", "world")
	fmt.Println(result)
}
