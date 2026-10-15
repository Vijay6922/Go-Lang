package main

import "fmt"

func main () {

	matrixs:=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println("original matrix:",matrixs)
	temp:=[][]int{}
	for i:=0;i<len(matrixs);i++{
		row:=[]int{}
		for j:=0;j<len(matrixs[0]);j++{
			if i>=j {
				row=append(row,matrixs[i][j])
			} else {
				row = append(row, 0)
			}
			
		}
		temp = append(temp, row)
	}
	fmt.Println(temp)
}