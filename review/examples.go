package main

import (
	"fmt"
	"strings"
)

var y int //package scope

type person struct {
	fname string //mala litera nie bedzie widoczny po za pakietem
	lname string
}

type secretAgent struct {
	person
	hasGun bool
}

//method
func (p person) allUpper() string {
	return strings.ToUpper(p.fname + " " + p.lname)
}
func (p secretAgent) allUpper() string {
	return strings.ToUpper("Agent specjalny " + p.fname + " " + p.lname)
}

func (p person) speak() {
	fmt.Println("mowi ", p.fname, p.lname)
}
func (p secretAgent) speak() {
	fmt.Println("mowi agent ", p.fname, p.lname)
}

//interface
//zmienna moze miec wiecej niz jeden typ, person bedzie tez human
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	// variables
	x := 1 //block scope

	//composed literal s type{data}
	xi := []int{1, 2, 3}
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	p1 := person{
		fname: "Lukasz",
		lname: "Wysocki",
	}

	sa1 := secretAgent{
		p1,
		true,
	}

	sa2 := secretAgent{
		person{
			"Jan",
			"Kowalski",
		},
		false,
	}

	fmt.Println("Hello, playground", x+y)
	fmt.Printf("%T", x+y)
	fmt.Printf("%T", xi)
	fmt.Printf("%T\n", m)
	fmt.Println(p1)
	fmt.Println(p1.allUpper())
	fmt.Println(sa1)
	fmt.Println(sa1.allUpper())
	fmt.Println(sa2.allUpper())
	fmt.Println(sa2.person.allUpper())

	//polimorfizm
	(saySomething(p1))
	(saySomething(sa2))

}

//funckcje
//func (reciver) nazwa(params) (retuen) {code}
