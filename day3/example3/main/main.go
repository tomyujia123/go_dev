package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "   hello world abc   \n"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("replace:", result)

	count := strings.Count(str, "l")
	fmt.Println("count:", count)

	result = strings.Repeat(str, 3)
	fmt.Println("repeat", result)

	result = strings.ToLower(str)
	fmt.Println("tolower:", result)

	result = strings.ToUpper(str)
	fmt.Println("toupper:", result)

	result = strings.TrimSpace(str)
	fmt.Println("trimspace:", result)

	result = strings.Trim(str, " \n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(str, " \n\r")
	fmt.Println("trimleft:", result)

	result = strings.TrimRight(str, " \n\r")
	fmt.Println("trimright:", result)

	splitResult := strings.Fields(str)
	fmt.Println("fileds:\n")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str, "l")
	fmt.Println("split\n")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	str2 = strconv.Itoa(100)
	fmt.Printf("str2 = %q", str2)

	num, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("can not convert")
	} else {
		fmt.Printf("num = %d", num)
	}
}
