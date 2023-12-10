package main

import "fmt"

type Matrix [][]int

func validationSequence(sudoArr []int) int {
	for i := 0; i < len(sudoArr)-1; i++ {
		for j := i; j < len(sudoArr); j++ {
			if sudoArr[i] > sudoArr[j] {
				sudoArr[i], sudoArr[j] = sudoArr[j], sudoArr[i]
			}
		}
	}
	valid := 1
	for i := 0; i < len(sudoArr) && valid == 1; i++ {
		if sudoArr[i] != i+1 {
			valid = 0
		}
	}
	return valid
}

func (sudo Matrix) checkCube() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var completed []int
			for row := i * 3; row < i*3+3; row++ {
				for col := j * 3; col < j*3+3; col++ {
					completed = append(completed, sudo[row][col])
				}
			}
			fmt.Println(completed)
			checkForValid(completed)
		}
	}
}

func (sudo Matrix) checkRows() {
	for i := 0; i < 9; i++ {
		var completed []int
		for j := 0; j < 9; j++ {
			completed = append(completed, sudo[i][j])
		}
		fmt.Println(completed)
		checkForValid(completed)
	}
}

func (sudo Matrix) checkColumns() {
	for i := 0; i < 9; i++ {
		var completed []int
		for j := 0; j < 9; j++ {
			completed = append(completed, sudo[j][i])
		}
		fmt.Println(completed)
		checkForValid(completed)
	}
}

func checkForValid(completed []int) {
	if validationSequence(completed) == 1 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func main() {
	sudoku := Matrix{
		{6, 4, 2, 1, 8, 3, 5, 9, 7},
		{3, 1, 9, 5, 2, 7, 6, 8, 4},
		{5, 8, 7, 9, 4, 6, 3, 2, 1},
		{8, 9, 4, 2, 6, 5, 7, 1, 3},
		{7, 3, 6, 8, 1, 9, 2, 4, 5},
		{2, 5, 1, 7, 3, 4, 9, 6, 8},
		{4, 6, 5, 3, 9, 1, 8, 7, 2},
		{1, 7, 8, 6, 5, 2, 4, 3, 9},
		{9, 2, 3, 4, 7, 8, 1, 5, 6},
	}
	fmt.Println("checking cubes")
	sudoku.checkCube()
	fmt.Println("checking rows")
	sudoku.checkRows()
	fmt.Println("checking columns")
	sudoku.checkColumns()
}
