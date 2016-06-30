package main

import "fmt"

var (
	a int
	b int
	c int
	n int
)

func main() {
	//
	//
	//
	// Знайти всi трiйки цiлих чисел a, b, c, що не перевищують n i задовiльняють
	// рiвняння a^2 + b^2 = c^2 .
	//
	//
	n = 100                  //Зададим наше n
	for i := 1; i < n; i++ { // цикл от 0 до 9, перебор всех возможных чисел,
		//далее 3 цикла, в каждом инициализируем переменную a,b,c
		a = i
		for j := 1; j < n; j++ {
			b = j
			for k := 1; k < n; k++ {
				c = k
				if twoNumSqrt(a, b) == sqrt(c) {
					fmt.Print("a = ", a, " b = ", b, " c = ", c, "\n")   //напишет переменные a,b,c
					fmt.Print(sqrt(a), "+", sqrt(b), "=", sqrt(c), "\n") //выведет нормальное уравнение
				}
			}
		}
	}

}

func sqrt(number int) int {
	numberSqrt := number * number // все что делает функция - тупо возносит в квадрат
	return numberSqrt
}

func twoNumSqrt(a, b int) int {
	cSqrtEqual := sqrt(a) + sqrt(b) // уравнение вернет квадрат двух чисел
	return cSqrtEqual
}
