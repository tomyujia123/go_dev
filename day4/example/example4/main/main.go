package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

}
