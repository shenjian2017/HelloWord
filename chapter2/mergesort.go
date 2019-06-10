package main

import "fmt"

func merge (A []int, B []int) []int {
	fmt.Println("merge: A>>", A, "B>>", B)
	var i,j,k int
	l := len(A) + len(B)
	C := make([]int, l, l)
	for (i<len(A) && j<len(B)) {
		if A[i] <= B[j] {
			C[k] = A[i]
			i++
		}else{
			C[k] = B[j]
			j++
		}
		k++
	}
	if i != len(A) {
		for ;i<len(A);i++ {
			C[k] = A[i]
			k++
		}
	}else if j != len(B) {
		for ;j<len(B);j++ {
			C[k] = B[j]
			k++
		}
	}
	
	fmt.Println("merge: C>>", C)
	return C
}

//(2.3)
func mergeSort (Z []int) []int {
	fmt.Println("mergeSort:Z>>", Z)
	l := len(Z)
	if l > 1 {
		A := Z[:l/2]
		B := Z[l/2:]
		AA := mergeSort(A)
		BB := mergeSort(B)
		return merge(AA, BB)
	}else{
		return Z
	}
}

func main () {
	LL := mergeSort([]int{2,4,5,7,1,2,3,6,8})
	fmt.Println("Output:", LL)
}
