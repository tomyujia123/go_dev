package main

import (
	"fmt"
	"time"
)

const (
	Female = 2
	Man    = 1
)

func main() {
	for {
		second := time.Now().Unix()
		if second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("male")
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
