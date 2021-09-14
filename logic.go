package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var judgePairs = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

func judge() string {
	for _, v := range judgePairs {
		if state[v[0]] == state[v[1]] && state[v[1]] == state[v[2]] {
			return state[v[0]]
		}
	}
	return ""
}

func enemyInput() {
	validCell := []int{}
	for i, v := range state {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		if intValue == i+1 {
			validCell = append(validCell, i)
		}
	}

	modCell := validCell[rand.Intn(len(validCell))]
	state[modCell] = ""
	fmt.Print("\n")
}
