package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80

	fmt.Println(stu)

	var stu1 *Student = &Student{
		Age:  20,
		Name: "hua",
	}

	fmt.Println(stu1.Name)

}
