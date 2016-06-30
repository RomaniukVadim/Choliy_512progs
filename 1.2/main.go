package main

import "fmt"

func main() {
	n := 3

	// задаем нужное количество строк
	twoD := make([][]int, n) // делаем двумерный срез
	for i := 0; i < n; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
	}
	twoD[0][0] = 1
	twoD[1][0] = 1
	twoD[1][1] = 1
	for i := 2; i < n; i++ {
		for j := 0; j < n+1; j++ {
			if (j == 0) || (j == n+1) {
				twoD[i][j] = 1
			} else {
				twoD[i][j] = twoD[i-1][j-1] + twoD[i-1][j]
			}
		}
	}
	for j := 0; j < n+1; j++ {
		if twoD[n][j] != 0 {
			fmt.Print("%d ", twoD[n][j])
		}
	}

}
