package main

import "fmt"

//bubble sort(2-2)
func bubbleSort (A []int) []int {
	var i,j int
	l := len(A)
	fmt.Println("bubbleSort input: ", A)
	for i = 0; i < l-1; i++ {
		for j = 0; j < l - i - 1; j++ {
			if A[j] > A[j+1] {
				A[j],A[j+1] = A[j+1],A[j]
			}
			fmt.Println("bubbleSort step", i, j, A)
		}
		fmt.Println("bubbleSort over", i, A)
	}

	return A
}


func main () {
	a := bubbleSort([]int{5,2,4,6,1,3})
	fmt.Println("Output: ", a)
}
