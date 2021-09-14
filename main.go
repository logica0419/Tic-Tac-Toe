package main

import (
	"fmt"
)

var state [9]string

func main() {
	displayInit()

	for {
		state = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		displayCurrent()

		for i := 0; i < 9; i++ {
			playerInput()
			displayCurrent()

			res := judge()
			if res != "" {

				break
			}
		}

		displayRes(judge())

		if ifRestart() {
			displayRestart()
		} else {
			break
		}
	}

	fmt.Print("Bye\n")
}
