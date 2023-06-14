package main

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/
type Employee interface {
	FullName()
	// добавим
	Accept(Visitor)
}
type Developer struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Developer) FullName() {
	fmt.Println("Developer ", d.FirstName, " ", d.LastName)
}

type Director struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Director) FullName() {
	fmt.Println("Director ", d.FirstName, " ", d.LastName)
}

// расширение интерфейса
type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}

type CalculIncome struct {
	bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}
func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}
func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}

func main() {
	backend := Developer{"Bob", "Bilbo", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}
	backend.FullName()
	backend.Accept(CalculIncome{20})
	boss.FullName()
	boss.Accept(CalculIncome{10})
}
