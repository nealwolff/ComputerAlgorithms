/*
Neal Wolff
Find Max Sub Array
Programed in go Language
*/

package main

import (
	"fmt"
	"time"
	)

func main() {
	A := []int {0,-81,0,17,-8,-20,4,74,40,28,65,-3,45,82,68,9,41,70,-54,15}
	start := time.Now()
	maxarr,l,h := maxSubArray(A,0,len(A)-1)
	elapsed := time.Since(start)

	fmt.Println("Low: ",A[l], " at index", l, "\nHigh: ", A[h],"at index", h, "\nMax", maxarr)
	fmt.Println("time taken is: ", elapsed, "\n")

	start = time.Now()
	fmt.Println("Using Kadane's algorithm, Max Sum: ", KmaxSub(A,len(A)))
	elapsed = time.Since(start)
	fmt.Println("time taken is: ", elapsed, "\n")
}

//function finds the maximum sub array using Divide and conquer
//three return values, the maximum number and the low and high numbers of the max sum subarray
func maxSubArray(arr [] int,l int, h int)(int, int, int){
	//base case
	//if arr recursively reaches a single value, return it.
	if l == h{
		return arr[l],l,l
	}
 
	//center of the array
	m := (l+h)/2
	
	maxLeftHalf, lhl, lhh := maxSubArray(arr, l, m)
	maxRighthalf, rhl,rhh := maxSubArray(arr, m+1, h)
	maxCrossMid, cml, cmh := maxCross(arr, l, m, h)
	
	maxOfAll := maxThree(maxLeftHalf, maxRighthalf, maxCrossMid)


	retLeft :=0
	retRight :=0
	//deterimes which right or left value to send on
	if maxOfAll == maxLeftHalf{
		retLeft = lhl
		retRight = lhh
	}else if maxOfAll == maxRighthalf{
		retLeft = rhl
		retRight = rhh
	}else if maxOfAll == maxCrossMid{
		retLeft = cml
		retRight = cmh
	}

	return maxOfAll,retLeft,retRight;
}

//This function finds the sum in the cross subarray
func maxCross(arr [] int,l int,m int, h int)(int, int, int){
	sum := 0
	lr :=0 //low return
	rr := 0 //high return

	leftSum := -999999

	//from the center, move backwards and add the values
	//the low return will be the value of i
	for i := m; i >= l; i--{
		sum = sum + arr[i]
		if sum > leftSum{
			leftSum = sum
			lr=i
		}
	}
 
	
	sum = 0
	rightSum := -999999

	//from the center, move upwards
	//the high return wil be the value of i
	for i := m+1; i <= h; i++{
		sum = sum + arr[i]
		if sum > rightSum{
			rightSum = sum
			rr=i
		}
	}
 
	//find the sum of the elements on the right and left of the midpoint
	sumofboth :=leftSum + rightSum
	return sumofboth, lr, rr
}
//function finds the max between three numbers
func maxThree(a int, b int, c int) int{ 
	maxInt:=0
	if a > b{
		maxInt = a
	}else{
		maxInt = b
	}

	if maxInt < c{
		maxInt =c
	}
	return maxInt
}

//Kadane's algorithm with time complexity O(n)
func KmaxSub(arr[] int, n int) int{
	sum := 0
	ret := 0
	
	for i :=0;i < n; i++{
		if sum+arr[i] > 0{
			sum+=arr[i]
		}else{
			sum  = 0
		}
		ret = max(ret,sum)	
	}
	return ret
}
func max(a int, b int) int{
	if a>b{
		return a
	}
	return b;
}
