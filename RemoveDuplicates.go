//https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import "fmt"

func main() {
	ary := []int{1, 1, 2,3,4,5,6,6,7,7}
	var nde = 0
	for i := 1; i < len(ary); i++ {
		if ary[nde] != ary[i] {
			ary[nde+1] = ary[i]
			nde++
		}
	}
	fmt.Print(nde+1)

}
