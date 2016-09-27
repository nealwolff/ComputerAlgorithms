/*
Neal Wolff
Merge Sort
Programed in go Language
*/

package main

import (
	"fmt"
	)

func main() {
	A := []int {0,17,2,16,7,18,19,6,3,4,1,20,5,14,12,15,10,9,13,8,11}
	mergeSort(A,1,len(A)-1)
	fmt.Println(A)
}
func mergeSort(A []int, p int, r int){
	if p<r{
         q := (p+r)/2
         mergeSort(A,p,q)
         mergeSort(A,q+1,r)
         merge(A,p,q,r)
	}
}

func merge (A []int,p int,q int,r int){
	n1 := q-p+1
	n2 := r-q
	L:= make([]int,n1+2)
	R:= make([]int,n2+2)
	i:=1
	j:=1

	for i=1;i<=n1;i++{
		L[i]=A[p+i-1]
	}
	for j=1;j<=n2;j++{
		R[j]=A[q+j]
	}
	largest:=FindLargest(A)+10
	L[n1+1]=largest
	R[n2+1]=largest

	i=1
	j=1

	for k:=p;k<=r;k++{
		if L[i]<=R[j]{
			A[k]=L[i]
			i=i+1
		}else{
 			A[k]=R[j]
			j=j+1
		}
	}
}

//finds the largest number in the array
func FindLargest(A[]int)int{
	largest := A[0]
	for i:=1;i<len(A);i++{
		if A[i] > largest{
			largest = A[i]
		}
	}
	return largest
}
