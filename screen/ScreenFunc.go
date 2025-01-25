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
	fmt.Print("\033[H") //указатель на начальное положение
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

func StartMenu(MenuValue *uint8) uint8 {
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
	fmt.Scan(&MenuValue)
	//time.Sleep(100 * time.Second)
	return *MenuValue
}

func ChoiceDifficulty(Difficulty *uint8) uint8 {
	fmt.Print("\033[H") //указатель на начальное положение
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
	fmt.Print("\033[H")   //указатель на начальное положение
	fmt.Print("\033[25B") //вниз
	fmt.Print("\033[18A") //вверх
	fmt.Print("\033[67C") //Вправо
	fmt.Print("ВЫберите сложность:")
	fmt.Print("\033[1B") //Вниз
	goToStartLn()
	fmt.Print("\033[67C") //Вправо
	fmt.Print("⟨1⟩Очень легкая")
	fmt.Print("\033[1B") //Вниз
	goToStartLn()
	fmt.Print("\033[67C") //Вправо
	fmt.Print("⟨2⟩Легкая")
	fmt.Print("\033[1B") //Вниз
	goToStartLn()
	fmt.Print("\033[67C") //Вправо
	fmt.Print("⟨3⟩Средняя")
	fmt.Print("\033[1B") //Вниз
	goToStartLn()
	fmt.Print("\033[67C") //Вправо
	fmt.Print("⟨3⟩Сложная")
	fmt.Print("\033[1B") //Вниз
	goToStartLn()
	fmt.Print("\033[67C") //Вправо
	fmt.Print("⟨4⟩Очень сложно")
	fmt.Print("\033[11B") //Вниз
	goToStartLn()
	fmt.Scan(&Difficulty)
	return *Difficulty
}
