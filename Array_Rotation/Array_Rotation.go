package main

import (
	"fmt"
	"os"
	"strconv"
)

func Solution(A []int, K int) []int {
	var temporary_val int           // переменная для хранения временного значения
	for i := 0; i < K%len(A); i++ { // крутим внешний цикл на величину сдвига
		temporary_val = A[len(A)-1]       // запоминаем послений элемент
		for j := len(A) - 1; j > 0; j-- { // крутим вложенный цикл для перемещения элементов по массиву
			A[j] = A[j-1]
		}
		A[0] = temporary_val // заносим в 0 элемент значение, которое запомнили
	}

	return A
}

func main() {
	var (
		N, K  int   // длина массива и величина сдвига
		array []int // исходный массив
	)

	fmt.Print("Введите N в диапазоне [0..100]: ") // ввод длины массива
	fmt.Scan(&N)

	if N < 0 || N > 100 { // проверка на корректный ввод
		fmt.Print("Некорректный ввод!")
		os.Exit(0)
	}

	array = make([]int, N, 100) // инициализация массива

	for i := 0; i < N; i++ {
		fmt.Print("Введите " + strconv.Itoa(i) + " элемент в диапазоне [-1000..1000]: : ")
		fmt.Scan(&array[i]) // ввод элементов массива

		if array[i] < -1000 || array[i] > 1000 { // проверка на корректный ввод
			fmt.Print("Некорректный ввод!")
			os.Exit(0)
		}
	}

	fmt.Print("Введите K в диапазоне [0..100]: ") // ввод величины сдвига
	fmt.Scan(&K)

	if K < 0 || K > 100 { // проверка на корректный ввод
		fmt.Print("Некорректный ввод!")
		os.Exit(0)
	}

	fmt.Println(array) // вывод старого массива

	array = Solution(array, K)

	fmt.Println(array) // вывод нового массива
}
