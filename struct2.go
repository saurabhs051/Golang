package main

import "fmt"

type Student struct {
	roll  uint16
	name  string
	marks uint16
}

func (s *Student) update(m uint16) { //pass by reference //i.e pointer receiver
	s.marks = m
}

func main() {
	s1 := Student{391, "Saurabh Singh", 58}
	fmt.Println("Marks : ", s1.marks)
	s1.update(98)
	fmt.Println("New Marks : ", s1.marks)
}
