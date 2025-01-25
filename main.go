package main

import (
	"Game/screen"
	"fmt"
	"os"
)

// Общая стркуктура для игрока и дилера
type Entity struct {
	health    uint8
	alive     bool
	inventory [8]uint8
	blocked   bool //Для наручников
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

// Для выбора сложности
var Difficulty uint8

// Для меню
var MenuValue uint8

func checkMenuValue(MenuValue uint8) uint8 {
check:
	fmt.Scan(&MenuValue)
	if MenuValue == 1 {
		return 1
	}
	if MenuValue == 2 {
		os.Exit(0)
	}
	if MenuValue > 2 {
		fmt.Print("\033[1A") //вверх
		fmt.Print("Введите число возле того действия, которое вы хотите совершить: ")
		goto check
	}
	return MenuValue
}

func main() {
	screen.InitScreenWithText()
	screen.StartMenu(&MenuValue)
	checkMenuValue(MenuValue)
	screen.ChoiceDifficulty(&Difficulty)
}
