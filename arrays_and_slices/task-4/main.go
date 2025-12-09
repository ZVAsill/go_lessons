package main

import "fmt"

/*
Задание 4. Матрица 3×3 и диагонали
Описание:Создай двумерный слайс matrix := [][]int размером 3×3.Заполни его числами по формуле matrix[i][j] = i+j.

Напиши две функции:
1. mainDiagonal(matrix [][]int) []int
2. secondaryDiagonal(matrix [][]int) []int
которые возвращают элементы главной и побочной диагонали.
Пример:
[ [0 1 2]  [1 2 3]  [2 3 4] ]Главная диагональ: [0 2 4]Побочная диагональ: [2 2 2]
*/
var matrixSize int = 3

func main() {
	matrix := make([][]int, matrixSize)
	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println(matrix)

	matrixMainDiagonal := make([]int, matrixSize)
	matrixMainDiagonal = mainDiagonal(matrix)
	fmt.Println(matrixMainDiagonal)

	matrixSecondaryDiagonal := make([]int, matrixSize)
	matrixSecondaryDiagonal = secondaryDiagonal(matrix)
	fmt.Println(matrixSecondaryDiagonal)

}

func mainDiagonal(matrix [][]int) []int {
	mmd := make([]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		mmd[i] = matrix[i][i]
	}
	return mmd
}

func secondaryDiagonal(matrix [][]int) []int {
	msd := make([]int, matrixSize)
	for i := 0; i < matrixSize; i++ {

		msd[i] = matrix[matrixSize-1-i][i]
	}
	return msd
}
