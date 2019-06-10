package main

import "fmt"

//incremental approach(2.1)
func insertSort1 (a []int)(b []int) {
	var i, j int
	fmt.Println("insertSort1 input>", a)
	for i=1; i < len(a); i++ {
		v := a[i]
		for j = i-1; j>=0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
				fmt.Println("insertSort1 L>", a)
			}else{
				break
			}
		}
		a[j+1]=v
		fmt.Println("insertSort1 I>", a)
	}

	//fmt.Println(a)
	return a
}


//recursive approach(2.3-4)
func insertSort2 (a []int, k int) []int{
	fmt.Println("insertSort2:", a)
	b := a[:k]
	if k > 0 {
		q := k - 1
		b = insertSort2(b, q)
		b = unit_sort(b, q, a[q])
		fmt.Println("insertSort2>", b)
	}

	return b
}

func unit_sort (a []int, n int, v int) ([]int){
	fmt.Println("unit_sort:", a, "n=", n, "v=", v)
	var i int
	b := make([]int, n+1)
	for i = 0; i < n; i++ {
		if  a[i] < v{
			b[i] = a[i]
		}else{
			b[i] = v
			c := b[:i+1]
			s := a[i:n]
			b = append(c, s...)
			fmt.Println("unit_sort>", b)
			return b
		}
	}
	b[i] = v
	fmt.Println("unit_sort>", b)
	return b
}

func main () {

	a := insertSort1([]int{5,2,4,6,1,3})
	fmt.Println("insertSort1 Output: ", a)
	fmt.Println()
	//a := unit_sort([]int{1,3,4,7,8}, 2, 6)
	b := insertSort2([]int{5,2,4,6,1,3}, 6)
	fmt.Println("insertSort2 Output: ", b)
}
