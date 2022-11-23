package main

import (
	"fmt"
	"os"
	"strconv"
)

func BubbleSort(ar []uint32) { // сортировка пузырьком
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, uint32(i), uint32(j))
			}
		}
	}
}

func swap(ar []uint32, i, j uint32) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

func Solution(A []uint32) uint32 {
	var result uint32 = 1 // переменная для возвращения значения метода Solution
	BubbleSort(A)         //сортировка пузырьком

	for i := 0; i < len(A)-1; i++ { // проверка последовательности
		if A[i] != uint32(i+1) { // если последовательность нарушена, то возвращаем 0
			result = 0
			return result
		}
	}
	return result
}

func main() {
	var Enter_length uint32 // строка для хранения длины массива
	var array []uint32      // массив чисел

	fmt.Print("Введите длину массива диапазоне [1..100000]: ") // ввод длины массива
	fmt.Scan(&Enter_length)

	if Enter_length != 0 { //проверяем "на дурака", если некорректно введена длина, завершаем программу
		array = make([]uint32, Enter_length, 100000) // устанавливаем длину массива

		for i := 0; i < len(array); i++ { // ввод массива с клавиатуры
			fmt.Print("Введите " + strconv.Itoa(i) + " элемент: ")
			fmt.Scan(&array[i])
			if array[i] <= 0 {
				fmt.Println("Некорректный ввод!") //если пользователь не умеет читать
				os.Exit(0)
			}
		}

		fmt.Println(Solution(array)) // вывод на экран является ли массив последовательностью

	} else {
		fmt.Println("Некорректный ввод!") //если пользователь не умеет читать
		os.Exit(0)
	}
}
