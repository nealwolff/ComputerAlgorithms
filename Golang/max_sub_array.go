/*
Neal Wolff
Maximum sum sub array CVS reader 1.0
Programed in go Language
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
	"time"
	"math/rand"
)

func main() {
		fmt.Println("\n|---------------------------------------|")
		fmt.Println("|WELCOME TO THE CVS MAX SUM ARRAY READER|")
		fmt.Println("|---------------------------------------|")
		for{
		fmt.Println("\nWould you like to read from the cvs or enter your own array?")
		fmt.Println("type 1 to open the csv file\nOR type 0 to enter the size for a random array\nOR type 2 to exit")
		

		var inp string
    		fmt.Scanln(&inp)
		in,err :=strconv.Atoi(inp)
		if err != nil{
			fmt.Println("\ninvalid input")
			in=4;
		}
		if in >2 || in<0{
			continue
		}
		if in ==2{
			break
		}
		if in == 1{
	
			//Finds the number of numbers in a file
			count, largestNum :=numOfLines("MaxSubArray_09.csv")
	
			//array to hold all the values
			array := make([] int, count)
			i :=0;
			f, _ := os.Open("MaxSubArray_09.csv")
			scanner := bufio.NewScanner(f)
	
			scanner.Split(bufio.ScanWords)

			// Scan all the lined from the file into the array
			for scanner.Scan() {
				line := scanner.Text()
				line = strings.Replace(line, ",", "", -1)
				array[i],_ = strconv.Atoi(line)
				i++
			}

			//find the subarray with the maximum sum
	
			start := time.Now()
			maxarr,l,h := maxSubArray(array,0,len(array)-1)
			elapsed := time.Since(start)

			fmt.Println("\nCSCI 5330 Spring 2016\nNeal Wolff\nRunning data for file\nMaxSubArray_09.csv\n")
			fmt.Println("Low: ",array[l], " at index", l, "\nHigh: ", array[h],"at index", h, "\nMax", maxarr, "\nmax in file", largestNum)
			fmt.Println("time taken is: ", elapsed, "\n")
	
			start = time.Now()
			fmt.Println("Using Kadane's algorithm, Max Sum: ", KmaxSub(array,len(array)))
			elapsed = time.Since(start)
			fmt.Println("time taken is: ", elapsed, "\n")
		}else{
			
			in2 :=3
			for{
				fmt.Println("Enter the number of random Integers to be writen to ARRAY.TXT")
		    		var inp2 string
    				fmt.Scanln(&inp2)
				in3,err :=strconv.Atoi(inp2)
				if err != nil{
					fmt.Println("\ninvalid input")
					continue
				}
				in2 =in3;
				break
			}

			
			RandArray := createRandArray(in2);
			filename := "ARRAY.TXT"
 
  			fmt.Println("writing: " + filename)
  			wf, _ := os.Create(filename)
			for i:=0;i<in2;i++{
				s := strconv.Itoa(RandArray[i])+"\n"
				io.WriteString(wf,s)
			}
			wf.Close()
			//find the subarray with the maximum sum
	
			start := time.Now()
			maxarr,l,h := maxSubArray(RandArray,0,len(RandArray)-1)
			elapsed := time.Since(start)

			fmt.Println("\nCSCI 5330 Spring 2016\nNeal Wolff\nRunning data for Random Array\n")
			fmt.Println("Low: ",RandArray[l], " at index", l, "\nHigh: ", RandArray[h],"at index", h, "\nMax", maxarr)
			fmt.Println("time taken is: ", elapsed, "\n")
	
			start = time.Now()
			fmt.Println("Using Kadane's algorithm, Max Sum: ", KmaxSub(RandArray,len(RandArray)))
			elapsed = time.Since(start)
			fmt.Println("time taken is: ", elapsed, "\n")
		}
	}

}

//function finds the number of lines in the file, also returns the largest number for the lels
func numOfLines(input string)(int,int){
	f, _ := os.Open(input)
	scanner := bufio.NewScanner(f)
	count:=0
	scanner.Split(bufio.ScanWords)

	//pulls the first number from the text
	largestNum,_:= strconv.Atoi(strings.Replace(scanner.Text(), ",", "", -1))

	// Scan all numbers and find the count and largest number
	for scanner.Scan() {
		tempNum,_:= strconv.Atoi(strings.Replace(scanner.Text(), ",", "", -1))
		if tempNum > largestNum{
			largestNum = tempNum
		}
		count++
	}
	count++
	f.Close()
	return count, largestNum
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
//creates an array with random values from 0-100 with size of n
func createRandArray(n int) []int{
	ret := make([] int, n)
	ret[0]=0 //makes sure there is one positive value for Kadane's algorithm
	for i :=1; i< n; i++{
		ret[i]=randInt(-100,100)
	}
	return ret
}
//random number generator
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
