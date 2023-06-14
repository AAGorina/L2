package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

type tree struct {
	hight      int
	wood_color string
	seeds      string
}

type Builder struct {
	tree tree
}

func newBuilder() *Builder {
	return &Builder{
		tree: tree{},
	}
}
func (b *Builder) setHight(h int) *Builder {
	b.tree.hight = h
	return b
}

func (b *Builder) setWoodColor(c string) *Builder {
	b.tree.wood_color = c
	return b
}

func (b *Builder) setSeeds(s string) *Builder {
	b.tree.seeds = s
	return b
}

func (b *Builder) build() tree {
	return b.tree
}

func main() {
	maple := newBuilder()
	fmt.Println(maple.setHight(15).setSeeds("вертолетики").setWoodColor("Темное").build())
}
