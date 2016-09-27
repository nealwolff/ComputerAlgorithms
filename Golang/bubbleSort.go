/*
Neal Wolff
Bubble Sort
Programed in go Language
*/

package main

import (
	"fmt"
	)

func main() {
	A := []int {17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	Bubble(A)
	fmt.Println(A)
}

func Bubble(A []int){
	for i:=0; i<len(A); i++{
		for j := len(A)-1; j>i;j--{
			if A[j] < A[j-1]{
				A[j],A[j-1] = A[j-1],A[j]
			}
		}
	}
}
