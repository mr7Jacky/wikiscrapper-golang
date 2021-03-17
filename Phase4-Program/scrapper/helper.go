package main

import (
	"fmt"
	"math"
	"strings"
)

// getContents used to concatenate all strings in contents to an output string
func getContents(contents []string, start int, end int) (str string) {
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

// topNIndex obtains the index of top n highest scores in a given list
func topNIndex(n int, list []float64) []int {
	// Initialization
	var l, curIndex, tmpIndex int
	var topScores []float64
	var topIndex []int
	topScores = make([]float64, n)
	topIndex = make([]int, n)
	// if the length of list is short than the given top n
	// Provide all the item index in list
	l = len(list)
	if l < n {
		n = l
		for i := 0; i < l; i++ {
			topIndex[i] = i
		}
	} else {
		// Else find the top n index with top n highest scores
		var curScore, tmpScore float64
		for i := 0; i < l; i++ {
			curScore = list[i]
			curIndex = i
			for j := 0; j < n; j++ {
				if topScores[j] == curScore {
					break
				}
				if topScores[j] < curScore {
					// Assign score and index to temp value for swap
					tmpScore = topScores[j]
					tmpIndex = topIndex[j]
					// Assign the value to lists to save
					topScores[j] = curScore
					topIndex[j] = curIndex
					// Finish swap
					curScore = tmpScore
					curIndex = tmpIndex
				}
			}
		}
	}
	return topIndex
}
