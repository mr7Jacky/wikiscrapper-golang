package main

import (
	"fmt"
	"math"
	"strings"
)

func printContents(contents []string, start int, end int) (str string) {
	str = ""
	for i := start; i < end; i++ {
		str += contents[i]
	}
	return str
}

// LCS is used to scoring the longest common substring of two given strings
func LCS(first string, second string) int {
	m := len(first) + 1
	n := len(second) + 1
	scoreMap := make([][]int, m)
	for i := 0; i < m; i++ {
		scoreMap[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				scoreMap[i][j] = 0
			} else if first[i-1] == second[j-1] {
				scoreMap[i][j] = scoreMap[i-1][j-1] + 1
			} else {
				scoreMap[i][j] = int(math.Max(float64(scoreMap[i-1][j]), float64(scoreMap[i][j-1])))
			}
		}
	}
	return scoreMap[m-1][n-1]
}

// modifySigmoid is a revised sigmoid function fit in the program
func modifySigmoid(x int) float64 {
	return -100/(1+math.Exp(-0.05*float64(x-50))) + 100
}

// maxArg returns the argument corresponding to the max value in the list
func maxArg(list []float64) int {
	var maxNum float64
	var maxIdx, l int
	maxNum = float64(math.MinInt8)
	maxIdx = 0
	l = len(list)
	for i := 0; i < l; i++ {
		if list[i] > maxNum {
			maxNum = list[i]
			maxIdx = i
		}
	}
	return maxIdx
}

// reformatSearchKey is used for return the correct format of search key
// e.g. "Iron Man" becomes "Iron+Man"
func reformatSearchKey(key string) string {
	fmt.Println(key)
	var reformatKey string
	splitKey := strings.Split(key, " ")
	reformatKey = strings.Join(splitKey, "+")
	return reformatKey
}

func splitParagraph(paragraph string) {
	splitKey := strings.Split(paragraph, "\n")
	fmt.Println(splitKey)
}
