package main

import (
	"fmt"
)

func main() {
	words := []string{"w1", "w1", "w2", "w3","w4"}
	f1 := "w1"
	f2 := "w2"
	var first = -1
	var min = 0
	var second = -1
	for i := 0; i < len(words); i++ {
		if words[i] == f1 {
			first = i
		} 
		 if words[i] == f2 {
			second = i
		
		}
		if (first != -1 && second != -1){
			var distance = Abs(first-second)
			if min == 0 || min > distance{
				min = distance
			}

		}


	}
	fmt.Println(min)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
