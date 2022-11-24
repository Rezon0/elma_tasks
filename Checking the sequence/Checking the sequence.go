package main

import (
	"fmt"
	"os"
	"strconv"
)

func Solution(A []int) int {
	var (
		result int = 1 // переменная для вывода результата вычислений
		count  int     // счетчик для поиска элемента
	)

	for i := 1; i <= len(A); i++ {
		count = 0
		for j := 0; j < len(A); j++ {
			if A[j] == i { // если элемент есть, то смотрим следующий
				count++
			}
		}
		if count == 0 || count > 1 { // если элемента не было или он повторяется, то возвращаем значение 0
			result = 0
			return result
		}
	}
	return result // если не было изменений с result выше, то возвращаем значение 1
}

func main() {
	var (
		N     int   // длина массива и величина сдвига
		array []int // исходный массив
	)

	fmt.Print("Введите N в диапазоне [1..100000]: ") // ввод длины массива
	fmt.Scan(&N)

	if N < 1 || N > 100000 { // проверка на корректный ввод
		fmt.Print("Некорректный ввод!")
		os.Exit(0)
	}

	array = make([]int, N, 100000) // инициализация массива

	for i := 0; i < N; i++ { // ввод элементов массива
		fmt.Print("Введите " + strconv.Itoa(i) + " элемент в диапазоне [1..1000000000]: ")
		fmt.Scan(&array[i]) // ввод элементов массива

		if array[i] < 1 || array[i] > 1000000000 { // проверка на корректный ввод
			fmt.Print("Некорректный ввод!")
			os.Exit(0)
		}

	}

	fmt.Println(Solution(array))
}
