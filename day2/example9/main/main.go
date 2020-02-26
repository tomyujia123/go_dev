package main

import "fmt"

//拷贝了副本进去，如果直接用a=b是不行的
func swap1(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	first := 100
	second := 200
	//first, second = swap(first, second)
	first, second = second, first
	fmt.Printf("first is %d and second is %d", first, second)
}
