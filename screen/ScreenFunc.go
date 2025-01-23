package screen

import (
	"fmt"
	"time"
)

func InitScreen() {

	fmt.Print("⌈")
	fmt.Print("\033[150C") //Вправо
	fmt.Print("⌉")
	fmt.Print("\033[24B") //вниз
	fmt.Print("\033[1D")  //влево
	fmt.Print("⌋")
	fmt.Print("\033[153D") //влево
	fmt.Print("⌊")
}

func InitScreenWithText() {

	fmt.Print("⌈")
	fmt.Print("\033[150C") //Вправо
	fmt.Print("⌉")
	fmt.Print("\033[24B") //вниз
	fmt.Print("\033[1D")  //влево
	fmt.Print("⌋")
	fmt.Print("\033[153D") //влево
	fmt.Print("⌊")
	fmt.Print("\033[15A") //вверх
	fmt.Print("\033[45C") //Вправо
	fmt.Print("Для более удобной игры, выставьте размер текста по уровням в углах")
	fmt.Print("\033[153D") //Влево
	fmt.Print("\033[15B")  //вниз

	time.Sleep(10 * time.Second)
}
