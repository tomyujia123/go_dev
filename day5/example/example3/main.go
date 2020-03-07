package main

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func main() {
	var root *Student = new(Student)
	root.Age = 18
	root.Name = "hua"
	root.Score = 80

}
