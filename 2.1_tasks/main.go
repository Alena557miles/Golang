package main

import (
	"fmt"
)

// Виконати популярні задачі із масивом:

func main() {

	fmt.Println(`====================================================================== `)
	// Знайти другий мінімальний елемент масиву

	fmt.Println(`FIRST TASK| Find second min elem in array: `)
	s := []int{1, 58, -36, -2, -7, 5, 4, 14}
	fmt.Println(s)
	min1 := s[0]
	min2 := s[0]
	for i := 0; i < len(s); {
		if s[i] < min1 {
			min1 = s[i]
		}
		if s[i] > min1 && min1 < min2 && s[i] < min2 {
			min2 = s[i]
		}
		i++
	}
	fmt.Println(`1st min elementof array: `, min1)
	fmt.Println(`2st min elementof array: `, min2)

	fmt.Println(`====================================================================== `)

	// Знайти неповторні цілі числа в масиві

	fmt.Println(`SECOND TASK|Find unique integers in an array: `)
	ar := []int{100, 2, 4, -14, -29, 2, 5, 4, 5}
	fmt.Println(ar)
	// array := []int{}
	index := 0
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[i] == ar[j] {
				index += 1
			}
		}
		if index < 2 {
			// array = append(array, ar[i])
			fmt.Println(ar[i])
		}
		index = 0
	}
	// fmt.Println(`uniuqe array: `, array)
	fmt.Println(`====================================================================== `)

	// Зверніть масив
	fmt.Println(`THIRD TASK| Print turned array: `)
	arr := []int{2, 3, 4, 5, 0, -3, -2}
	arrFinal := []int{}
	fmt.Println(`Array: `, arr)
	for i := len(arr) - 1; i > 0; i-- {
		arrFinal = append(arrFinal, arr[i])
	}
	fmt.Println(`Final Array:`, arrFinal)
	fmt.Println(`====================================================================== `)

	// Видаліть зі масива дублюючі значення

	fmt.Println(`FOURTH TASK|Return an unique array: `)
	myArr := []int{100, 2, 4, -14, -29, 2, 5, 4, 5}
	fmt.Println(myArr)

	arrayUnique := []int{}
	ind := 0
	for i := 0; i < len(myArr); i++ {
		for j := 0; j < len(myArr); j++ {
			if ar[i] == ar[j] {
				ind += 1
			}
		}
		if ind < 2 {
			arrayUnique = append(arrayUnique, myArr[i])
		}
		ind = 0
	}
	fmt.Println(`unique array: `, arrayUnique)
	fmt.Println(`====================================================================== `)

	// Порахувати в рядку із латинських і кириличних симболів кількість перших і других
	fmt.Println(`FIFTH TASK|Return amount of symbols from the string: `)
	str := "Slava Україні !"
	fmt.Println(`Given string: `, str)
	fmt.Println(`Amount of symbols:`, len([]rune(str)))
	fmt.Println(`====================================================================== `)

}
