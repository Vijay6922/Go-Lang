package main 

import "fmt"

func main (){

	matrixs:=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println("original matrix:",matrixs)

	for i:=0;i<len(matrixs);i++{
		mainDaignoal:=i
		anti:=len(matrixs)-1-i

		temp:=matrixs[i][mainDaignoal]

		matrixs[i][mainDaignoal]= matrixs[i][anti]

		matrixs[i][anti]=temp

	}
	fmt.Println("after: ",matrixs)
}