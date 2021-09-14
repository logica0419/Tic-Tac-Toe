package main

import (
	"fmt"
	"math/rand"
)

var state [9]string

func main() {
	displayInit()

	for {
		diff := selectDiff()

		turn := rand.Intn(2)
		displayTurn(turn)

		state = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		displayCurrent()

		for i := 0; i < 9; i++ {
			if i%2 == turn {
				playerInput()
			} else {
				enemyInput(diff)
			}
			displayCurrent()

			res := judge()
			if res != "" {
				break
			}
		}

		displayRes(judge())

		if ifRestart() {
			displayReset()
		} else {
			break
		}
	}

	fmt.Print("Bye-bye. Thank you for playing!\n")
}
