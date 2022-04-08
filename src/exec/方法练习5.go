package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) Say() string {
	info := fmt.Sprintf("%v,%v,%v,%v,%v", s.name, s.gender, s.age, s.id, s.score)
	return info
}

func main() {
	var s = Student{
		name:   "Tom",
		gender: "男",
		age:    18,
		id:     1,
		score:  15.6,
	}
	fmt.Println((&s).Say())
}
