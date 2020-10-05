/*Find All Anagrams in a String: Given a string s and a non-empty string p, find all the start indices of p's anagrams in s. Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100. The order of output does not matter.
Example:
 Input: s: "cbaebabacd" p: "abc"
Output: [0, 6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc". The substring with start index = 6 is "bac", which is an anagram of "abc".
Source: https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputS := "cbaebabacd"
	inputP := "abc"
	indexAray := []int{}
	inputP = SortString(inputP)
	j := len(inputP)
	for i := 0; i < len(inputS); i++ {
		chars1 := inputS[i:j]
		s1 := []byte(chars1)
		sort.Slice(s1, func(i int, j int) bool { return s1[i] < s1[j] })
		if string(s1) == inputP {
			indexAray = append(indexAray, i)
		}
		if j > len(inputS)-1 {
			break
		}
		j++
	}
	fmt.Println(indexAray)

}

// sort string
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
