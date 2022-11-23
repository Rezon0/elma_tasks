package main

import (
	"fmt"
	"strconv"
)

func Solution(A []int) int {
	var result int // переменная для возвращения значения метода Solution
	var flag bool  // флаг для различия найденных элементов от не найденных

	for i := 1; i <= len(A)+1; i++ { // цикл перебора чисел от [1 до N+1]
		flag = false                  // сбрасываем флаг
		for j := 0; j < len(A); j++ { // цикл для перебора элементов массива
			result = i     // записываем значение в result, если не выйдем досрочно из цикла, значит этот элемент наш
			if A[j] == i { // проверяем есть ли такой элемент в массиве
				flag = true
				break
			}
		}
		if flag == false { // если во вложенном цикле не был найден элемент, то попадем в этот блок и выходим из метода
			return result
		}
	}
	return result
}

func main() {
	var Enter_length uint32 // строка для хранения длины массива
	var array []int         // массив чисел

	fmt.Print("Введите длину массива диапазоне [1..100000]: ") // ввод длины массива
	fmt.Scan(&Enter_length)

	if Enter_length != 0 { //проверяем "на дурака", если некорректно введена длина, завершаем программу
		array = make([]int, Enter_length, 100000) // устанавливаем длину массива

		for i := 0; i < len(array); i++ { // ввод массива с клавиатуры
			fmt.Print("Введите " + strconv.Itoa(i) + " элемент: ")
			fmt.Scan(&array[i])
		}

		fmt.Println(Solution(array)) // вывод на экран недостающего элемента

	} else {
		fmt.Println("Некорректный ввод!") //если пользователь не умеет читать
	}
}
