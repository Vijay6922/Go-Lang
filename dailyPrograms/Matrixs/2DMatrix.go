package main

import "fmt"

func add(a[][]int,b[][]int)[][]int{
	sum:=make([][]int,len(a))
	for i:=0;i<len(a);i++ {
		sum[i] = make([]int, len(a[i]))
		for j:=0;j<len(a[i]);j++ {
			sum[i][j] = a[i][j]+b[i][j]
		}
	}
	return sum
}

func mul(a[][]int,b[][]int)[][]int{

	sum:=make([][]int,len(a))
	if len(a) == 0 || len(a[0]) != len(b) {
		fmt.Println("cannot multiply: incompatible dimensions")
		return nil
	}
	for i:=0;i<len(a);i++{
		sum[i]=make([]int,len(b[0]))
	}
	for i:=0;i<len(a);i++ {
		for j:=0;j<len(b[0]);j++{
			for k:=0;k<len(b);k++{
				sum[i][j] += a[i][k]*b[k][j]
			}

		}
	}
	return sum
}

func muls(a [][]int, b [][]int) [][]int {
	// m = rows of a, n = columns of a
	m := len(a)    // Value is 2
	n := len(a[0]) // Value is 3

	// p = rows of b, q = columns of b
	p := len(b)    // Value is 3
	q := len(b[0]) // Value is 2 (This is len(b[0]))

	// Safety Check: n must equal p
	if n != p {
		fmt.Println("Cannot multiply: columns of A must equal rows of B")
		return nil
	}

	// Step 1: Create 'm' rows for the answer matrix (Using len(a))
	sum := make([][]int, m) // creates [ [] , [] ]

	// Step 2: Put 'q' columns into each row (Using len(b[0]))
	for i := 0; i < m; i++ {
		sum[i] = make([]int, q) // Allocates 2 columns per row
	}
	// Now 'sum' looks like: [[0, 0], [0, 0]] (A perfect 2x2 matrix!)

	// Step 3: Run the 3 loops to calculate
	for i := 0; i < m; i++ { // Loops through rows of A (m)
		for j := 0; j < q; j++ { // Loops through columns of B (q)
			for k := 0; k < n; k++ { // Loops through the shared dimension (n or p)
				sum[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return sum
}
func main() {

	mat := [][]int{{1, 2}, {3, 4}}
	mat2 := [][]int{{5, 6}, {7, 8}}
	result :=add(mat,mat2)
	fmt.Println("addtion of 2 matrixs:",result)
	result2 :=mul(mat,mat2)
	fmt.Println("multiplication of 2 matrixs:",result2)
	result21 :=muls(mat,mat2)
	fmt.Println("multiplication of 2 matrixs:",result21)
}