package screen

import (
	"fmt"
	"time"
)

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
	fmt.Print("\033[40C") //Вправо
	fmt.Print("Для более удобной игры, выставьте размер текста по уровням в углах")
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
	fmt.Print("\033[20B") //Вниз
	fmt.Scan(&MenuValue)
	//Визуально очищаем значение пользователя, что бы не мешало
	fmt.Print("\033[153D") //влево
	fmt.Print(" ")
	fmt.Print("\033[H") //указатель на начальное положение
	time.Sleep(10 * time.Second)
}
