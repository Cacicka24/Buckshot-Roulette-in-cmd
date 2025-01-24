package screen

import (
	"fmt"
	"os"
	"time"
)

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

func goToStartLn() {
	fmt.Print("\r")
}

func cleanScreen() {
	//чистим экран
	for i := 1; i <= 24; i++ {
		fmt.Print("\033[K") // Очистить текущую строку до конца
		fmt.Println()
	}
}

func InitScreen() {
	cleanScreen()
	fmt.Print("⌈")
	fmt.Print("\033[150C") //Вправо
	fmt.Print("⌉")
	fmt.Print("\033[24B") //вниз
	fmt.Print("\033[1D")  //влево
	fmt.Print("⌋")
	goToStartLn()
	fmt.Print("⌊")
	fmt.Print("\033[H") //указатель на начальное положение
}

func InitScreenWithText() {
	fmt.Print("\033[H") //указатель на начальное положение
	fmt.Print("⌈")
	fmt.Print("\033[150C") //Вправо
	fmt.Print("⌉")
	fmt.Print("\033[24B") //вниз
	fmt.Print("\033[1D")  //влево
	fmt.Print("⌋")
	fmt.Print("\033[153D") //влево
	fmt.Print("⌊")
	fmt.Print("\033[15A") //вверх
	fmt.Print("\033[35C") //Вправо
	fmt.Print("Для более удобной игры, выставьте размер текста по уровням в углах")
	goToStartLn()
	fmt.Print("\033[1B")  //вниз
	fmt.Print("\033[35C") //Вправо
	fmt.Print(" Для взаимодействия с игрой вводите числа возле действия, которые вы хотите совершить")
	fmt.Print("\033[H") //указатель на начальное положение
	time.Sleep(10 * time.Second)
}

func StartMenu(MenuValue uint8) {
	//чистим экран
	cleanScreen()
	fmt.Print("\033[H") //указатель на начальное положение
	fmt.Print("⌈")
	fmt.Print("\033[150C") //Вправо
	fmt.Print("⌉")
	fmt.Print("\033[24B") //вниз
	fmt.Print("\033[1D")  //влево
	fmt.Print("⌋")
	goToStartLn()
	fmt.Print("⌊")
	fmt.Print("\033[18A") //вверх
	fmt.Print("\033[67C") //Вправо
	fmt.Print("\033[1;32mМЕНЮ\033[0m")
	//перемещение на строчку вниз, для отрисовки 1 кнопки
	fmt.Print("\033[4D") //влево
	fmt.Print("\033[1B") //Вниз
	fmt.Print("⟨1⟩ Играть")
	//перемещение на строчку вниз, для отрисовки 2 кнопки
	fmt.Print("\033[10D") //влево
	fmt.Print("\033[1B")  //Вниз
	fmt.Print("⟨2⟩ Выход")
	goToStartLn()
	fmt.Print("\033[15B") //Вниз
	checkMenuValue(MenuValue)
	time.Sleep(10 * time.Second)
}
