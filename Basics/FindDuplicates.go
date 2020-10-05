package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	inputAry := []int{2, 3, 5, 1, 4}
	flag := findDuplicateMethod1(inputAry)
	fmt.Println("Found duplicates : ", flag)
	flag = findDuplicateMethod2(inputAry)
	fmt.Println("Found duplicates : ", flag)
	flag = findDuplicateMethod3(inputAry)
	fmt.Println("Found duplicates : ", flag)
	flag = findDuplicateMethod4(inputAry)
	fmt.Println("Found duplicates : ", flag)

}
func findDuplicateMethod1(inputAry []int) int {
	for i := 0; i < len(inputAry); i++ {
		for j := i + 1; j < len(inputAry); j++ {
			if inputAry[i] == inputAry[j] {
				return inputAry[i]
			}
		}
	}
	return math.MinInt32
}

func findDuplicateMethod2(inputAry []int) int {
	sort.Ints(inputAry)
	for i := 1; i < len(inputAry); i++ {
		if inputAry[i] == inputAry[i-1] {
			return inputAry[i]
		}
	}
	return math.MinInt32

}
func findDuplicateMethod3(inputAry []int) int {
	len := len(inputAry)
	fmt.Println(inputAry)
	auxArray := make([]bool, len)
	for i := 0; i < len; i++ {
		tmp := inputAry[i]
		if auxArray[tmp-1] == true {
			return tmp
		}
		auxArray[tmp-1] = true
	}

	return math.MinInt32

}

func findDuplicateMethod4(inputAry []int) int {
	for i := 0; i < len(inputAry); i++ {
		tmp := Abs(inputAry[i])
		if inputAry[tmp-1] < 0 {
			return tmp
		}
		inputAry[tmp-1] *= -1
	}
	return math.MinInt32
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
