/*
Neal Wolff
Insertion Sort
Programed in go Language
*/

package main

import (
	"fmt"
	)

func main() {
	A := []int {17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	insertionSort(A)
	fmt.Println(A)
}
func insertionSort(A []int){
	for j :=2;j<= len(A);j++{
			key := A[j-1]
			i := j-1
			for i > 0 && A[i-1]> key{
				A[i] = A[i-1]
				i = i-1
			}
		A[i] = key
		}
}
