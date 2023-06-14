package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

type Window struct {
	state State
}

type State interface {
	Action(w *Window)
}

type WinClosed struct {
}

type WinOpen struct {
}

func (wc *WinClosed) Action(w *Window) {
	fmt.Println("Window is closed")
	w.state = &WinOpen{}
	fmt.Println("Open the Window")
}

func (wo *WinOpen) Action(w *Window) {
	fmt.Println("Window is open")
	w.state = &WinClosed{}
	fmt.Println("Close the window")
}

func main() {
	win := &Window{&WinOpen{}}
	for {
		win.state.Action(win)
		time.Sleep(3 * time.Second)
	}
}
