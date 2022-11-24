package main

import (
	"fmt"
	"os"
	"strconv"
)

func Solution(A []int) int { // метод поиска элемента без пары
	var (
		flag bool // флаг для отметки элемента с парой
		val  int // переменная для хранения значения
	)
	for i := 0; i < len(A); i++ {
		val = A[i]   // записываем значение
		flag = false // сбрасываем флаг
		for j := 0; j < len(A); j++ {
			if val == A[j] && i != j { // если такой элемент есть, и это не тот же самый, то ищем следующий
				flag = true
				break
			}
		}
		if flag == false { // если элемент оказался без пары, попадаем в это условие и выходим
			break
		}
	}

	return val // возвращает элемент без пары
}

func main() {
	var (
		array []int // массив
		N     int   // длина массива
	)

	fmt.Println("Введите N в диапазоне [1..1000000]: ") // ввод длины массива
	fmt.Scan(&N)

	if N < 1 || N > 1000000 { // проверка на корректный ввод
		fmt.Print("Некорректный ввод!")
		os.Exit(0)
	}

	array = make([]int, N, 1000000) // инициализация массива

	for i := 0; i < N; i++ { // ввод элементов массива
		fmt.Print("Введите " + strconv.Itoa(i) + " элемент: ")
		fmt.Scan(&array[i])

		if array[i] < 1 || array[i] > 1000000000 { // проверка на корректный ввод
			fmt.Print("Некорректный ввод!")
			os.Exit(0)
		}
	}

	fmt.Println(Solution(array)) // вывод на экран элемента без пары
}
