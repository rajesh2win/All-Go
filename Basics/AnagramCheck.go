package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := "abca"
	s2 := "bcad"
	flag := checkAnagramMethod1(s1, s2)
	fmt.Println(flag)
	flag = checkAnagramMethod2(s1, s2)
	fmt.Println(flag)

}
func checkAnagramMethod1(s1 string, s2 string) bool {
	chars1 := []byte(s1)
	chars2 := []byte(s2)
	sort.Slice(chars1, func(i int, j int) bool { return chars1[i] < chars1[j] })
	sort.Slice(chars2, func(i int, j int) bool { return chars2[i] < chars2[j] })
	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			return false
		}

	}
	return true
}
func checkAnagramMethod2(s1 string, s2 string) bool {
	charArray := [26]int{}
	s1len := len(s1)
	if s1len != len(s2) {
		return false
	}
	for i := 0; i < s1len; i++ {
		tmp := s1[i] - 97 //Ascii value of a=97
		charArray[tmp] = charArray[tmp] + 1
	}
	for i := 0; i < s1len; i++ {
		tmp := s2[i] - 97
		if charArray[tmp] <= 0 {
			return false
		}
		charArray[tmp] = charArray[tmp] - 1

	}

	return true
}
