package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear ", p.name)
	fmt.Println("We have some vacancies for you")
	for _, v := range vacancies {
		fmt.Println(v)
	}
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}
