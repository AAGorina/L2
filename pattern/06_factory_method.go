package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

type Animal interface {
	talk()
}
type Cat struct {
}

type Dog struct {
}

func (c *Cat) talk() {
	fmt.Println("Meow")
}

func (d *Dog) talk() {
	fmt.Println("Wuf")
}

type AnimalFactory interface {
	BupAnimal() Animal
}

type CatFactory struct {
}

func (c *CatFactory) BupAnimal() Animal {
	return &Cat{}
}

type DogFactory struct {
}

func (d *DogFactory) BupAnimal() Animal {
	return &Dog{}
}

func main() {
	cf := &CatFactory{}
	cat := cf.BupAnimal()
	cat.talk()
	df := &DogFactory{}
	dog := df.BupAnimal()
	dog.talk()
}
