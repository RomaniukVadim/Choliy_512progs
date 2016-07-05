package main

import (
	"fmt"
)

const startNum = 2 //Константа, начинаем поиск простых чисел с нее(не менять)
var endNum = 100   // число которым заканчиваем поиск(менять на свое)

//как работает метод можно посмотреть тут
//https://upload.wikimedia.org/wikipedia/commons/b/b9/Sieve_of_Eratosthenes_animation.gif

func main() {
	arrayOfNumbers := make([]int, endNum-1)
	for i := 2; i <= endNum; i++ {
		arrayOfNumbers[i-2] = i
	}

	for j := 0; j < endNum/4+1; j++ {
		p := arrayOfNumbers[j]
		arrayLang := len(arrayOfNumbers)
		for i := 1; i < arrayLang-i+2; i++ {
			if arrayOfNumbers[i]%p == 0 && p != arrayOfNumbers[i] {
				fmt.Println(arrayOfNumbers[i], i)
				arrayOfNumbers = append(arrayOfNumbers[:i], arrayOfNumbers[i+1:]...)
			}

		}
	}
	fmt.Println(arrayOfNumbers[:len(arrayOfNumbers)-10])
}
