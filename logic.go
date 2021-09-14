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

func enemyInput(diff int) {
	notToLose := false
	if diff == 2 {
		notToLose = enemyNotToLose()
	}

	toWin := false
	if !notToLose && diff > 0 {
		toWin = enemyToWin()
	}

	if diff > 0 && state[4] != "○" && state[4] != "x" {
		state[4] = "x"
	}

	if !notToLose && !toWin {
		enemyRand()
	}

	fmt.Print("\n")
}

func enemyNotToLose() bool {
	for _, v := range judgePairs {
		if state[v[0]] == "○" && state[v[1]] == "○" && state[v[2]] != "x" {
			state[v[2]] = "x"
			return true
		}

		if state[v[1]] == "○" && state[v[2]] == "○" && state[v[0]] != "x" {
			state[v[0]] = "x"
			return true
		}

		if state[v[0]] == "○" && state[v[2]] == "○" && state[v[1]] != "x" {
			state[v[1]] = "x"
			return true
		}
	}

	return false
}

func enemyToWin() bool {
	for _, v := range judgePairs {
		if state[v[0]] == "x" && state[v[1]] == "x" && state[v[2]] != "○" {
			state[v[2]] = "x"
			return true
		}

		if state[v[1]] == "x" && state[v[2]] == "x" && state[v[0]] != "○" {
			state[v[0]] = "x"
			return true
		}

		if state[v[0]] == "x" && state[v[2]] == "x" && state[v[1]] != "○" {
			state[v[1]] = "x"
			return true
		}
	}

	return false
}

func enemyRand() {
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
	state[modCell] = "x"
}
