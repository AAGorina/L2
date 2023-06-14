package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

type Command interface {
	Execute()
}

type Calculator struct {
	value int
}

// concrete cpmmands

type SubCommand struct {
	calculator *Calculator
	value      int
}

type SumCommand struct {
	calculator *Calculator
	value      int
}

func (c *Calculator) Sum(value int) {
	c.value += value
}

func (c *Calculator) Sub(value int) {
	c.value -= value
}

func (s *SubCommand) Execute() {
	s.calculator.Sub(s.value)
}

func (s *SumCommand) Execute() {
	s.calculator.Sum(s.value)
}

// управление
type Invoke struct {
	commands []Command
}

func (i *Invoke) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoke) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	calc := &Calculator{}
	sum := &SumCommand{calculator: calc, value: 20}
	sub := &SubCommand{calculator: calc, value: 5}
	inv := &Invoke{}
	inv.StoreCommand(sum)
	inv.StoreCommand(sub)
	inv.ExecuteCommands()
	fmt.Println(calc.value)
}
