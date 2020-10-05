/*
Input: "algorithmica is your best guru" Output: " acimhtirogla si ruoy tseb urug"

1) Reverse Words in a String: Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order. Example:
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
Source: https://leetcode.com/problems/reverse-words-in-a-string-iii/ */

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "algorithmica is your best guru"
	reverse := ""
	for _, v := range strings.Fields(s) {
		result := ""
		for _, v := range v {
			result = string(v) + result
		}
		reverse += " " + result
	}

	fmt.Println(reverse)

}
