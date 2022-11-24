package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		temporary_val int   // переменная для хранения временного значения
		N, K          int   // длина массива и величина сдвига
		array         []int // исходный массив
	)

	fmt.Print("Введите N: ") // ввод длины массива
	fmt.Scan(&N)
	array = make([]int, N, 100) // инициализация массива

	for i := 0; i < N; i++ {
		fmt.Print("Введите " + strconv.Itoa(i) + " элемент: ")
		fmt.Scan(&array[i]) // ввод элементов массива
	}

	fmt.Print("Введите K: ") // ввод величины сдвига
	fmt.Scan(&K)

	fmt.Println(array) // вывод старого массива

	for i := 0; i < K%N; i++ { // крутим внешний цикл на величину сдвига
		temporary_val = array[N-1]   // запоминаем послений элемент
		for j := N - 1; j > 0; j-- { // крутим вложенный цикл для перемещения элементов по массиву
			array[j] = array[j-1]
		}
		array[0] = temporary_val // заносим в 0 элемент значение, которое запомнили
	}

	fmt.Println(array) // вывод нового массива
}
