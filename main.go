package main

import (
	"Game/screen"
)

// Общая стркуктура для игрока и дилера
type Entity struct {
	health    uint8
	alive     bool
	inventory [8]uint8
	blocked   bool
}

func (e *Entity) get_1_damage() {
	e.health = e.health - 1
}

func (e *Entity) get_2_damage() {
	e.health = e.health - 2
}

// Общая структура для предметов
type Item struct {
	name string
	ID   uint8
}

func main() {
	screen.InitScreenWithText()
}
