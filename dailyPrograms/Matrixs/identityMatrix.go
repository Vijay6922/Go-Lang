package main

import "fmt"

func main() {

	matrixs := [][]int{{1, 0}, {0, 1}}

	rows := len(matrixs)
	cols := len(matrixs[0])
	identity:=true

	// 1. First rule: An identity matrix MUST be a square
	if rows != cols {
		fmt.Println("Not an identity matrix (not square)")
		return
	}

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols; j++ {
			if i == j {
				if matrixs[i][j]!=1{
					identity=false
					break
				}
			} else {
				if matrixs[i][j]!=0 {
					identity=false
					break
				}
			}
		}
	}
	fmt.Println(identity)
}