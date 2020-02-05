package main

import (
	"fmt"
	a "go_dev/day2/example2/add"
)

func init()  {
	fmt.Println("initialized")
}

func main() {
	//a.Test()
	fmt.Println("Name =", a.Name)
	fmt.Println("age =", a.Age)
}
