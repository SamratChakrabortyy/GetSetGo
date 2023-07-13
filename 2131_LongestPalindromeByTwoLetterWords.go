package main

import (
	"fmt"
	"math"
)

func longestPalindrome(words []string) int {
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word] = countMap[word] + 1
	}
	//fmt.Println(words)
	result := 0
	central := false
	for word, count := range countMap {
		if word[0] == word[1] {
			if count%2 == 0 {
				result += 2
			} else {
				result += count - 1
				central = true
			}
		} else if word[0] < word[1] {
			revWord := string(word[1]) + string(word[0])
			result += int(math.Min(float64(count), float64(countMap[revWord])))
		}
		//fmt.Println(word, count, result, central)
	}

	if central {
		result++
	}
	return result * 2
}
