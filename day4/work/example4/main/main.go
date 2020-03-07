package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (wordCount, spaceCount, numCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("%d %d %d %d", wc, sc, nc, oc)
}
