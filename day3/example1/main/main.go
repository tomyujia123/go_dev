package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m)
	for i := n; i < m; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}
