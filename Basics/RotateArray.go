package main

import "fmt"

func main() {
	inputAry := []int{2, 3, 5, 7, 11, 13}
	k := 4
	rotateArray(inputAry, k)
	rotatWithAuxilaryArray(inputAry, k)
	rotatBySwappingTwoEnds(inputAry, k)

}

func rotateArray(inputAry []int, k int) {
	len := len(inputAry) - 1

	for i := 0; i < k; i++ {
		tmp := inputAry[len]
		for p := len; p >= 0; p-- {
			if p == 0 {
				inputAry[p] = tmp
			} else {
				inputAry[p] = inputAry[p-1]
			}
		}

	}
	fmt.Println(inputAry)
}

func rotatWithAuxilaryArray(inputAry []int, k int) {
	len := len(inputAry)
	auxArray := []int{}
	for i := 0; i < len; i++ {
		if i+k < len {
			auxArray[i+k] = inputAry[i]
		} else {
			auxArray[(i+k)%len] = inputAry[i]
		}
	}
	fmt.Println(auxArray)
}

func rotatBySwappingTwoEnds(inputAry []int, k int) {
	fmt.Println(inputAry)

	len := len(inputAry)
	reverse(inputAry, 0, len-1)
	reverse(inputAry, 0, k-1)
	reverse(inputAry, k, len-1)
	fmt.Println(inputAry)

}

func reverse(inputAry []int, l int, r int) {
	for l < r {
		tmp := inputAry[l]
		inputAry[l] = inputAry[r]
		inputAry[r] = tmp
		l++
		r--
	}

}
