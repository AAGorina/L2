package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

type MessageMethod interface {
	WriteMessage(str string)
}

type PaperLetter struct {
	adress string
}

type ELetter struct {
	email string
}

func (p *PaperLetter) WriteMessage(str string) {
	fmt.Println("You got message - ", str, " - from adress - ", p.adress)
}

func (e *ELetter) WriteMessage(str string) {
	fmt.Println("You got message - ", str, " - from email - ", e.email)
}

// context analog
type DeliverMessage struct {
	messagemethod MessageMethod
}

func (d *DeliverMessage) SetMessageMethod(mm MessageMethod) {
	d.messagemethod = mm
}

func (d *DeliverMessage) SendMessage(str string) {
	d.messagemethod.WriteMessage(str)
}

func main() {
	pl := &PaperLetter{"Представьте что тут адресс"}
	el := &ELetter{"адресс@gmail.com"}

	d := &DeliverMessage{}
	d.SetMessageMethod(pl)
	d.SendMessage("Привет")

	d.SetMessageMethod(el)
	d.SendMessage("Пока")
}