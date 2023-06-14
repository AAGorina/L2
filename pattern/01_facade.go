package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/
// процессор
type CPU struct{}

func (*CPU) Freeze() {
	fmt.Println("CPU Freeze")
}

func (*CPU) Jump(position int) {
	fmt.Printf("CPU Jump to %d\n", position)
}

func (*CPU) Execute() {
	fmt.Println("CPU Execute")
}

// память
type Memory struct{}

func (*Memory) Load(position int, data string) {
	fmt.Printf("Memory Load data '%s' to position %d\n", data, position)
}

// ЖД
type HardDrive struct{}

func (*HardDrive) Read(position int, size int) string {
	data := fmt.Sprintf("HardDrive Read data from position %d with size %d", position, size)
	fmt.Println(data)
	return data
}

// фасад компьютера который обьединяет их вместе
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// конструктор копьютера
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// при запуска настраивается проц и память
func (c *ComputerFacade) Start() {
	c.cpu.Freeze()
	c.memory.Load(0, "boot_loader")
	c.cpu.Jump(0)
	c.cpu.Execute()
}

func main() {
	computer := NewComputerFacade()
	// можно вызвать для запуска только один метод вместо списка операций
	computer.Start()
}
