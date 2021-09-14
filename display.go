package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func displayInit() {
	fmt.Print("\033[1m")
	fmt.Print("\033[4m")
	fmt.Print("Welcome to Tic-Tac-Toe!\n\n")
	fmt.Print("\033[0m")
}

func displayReset() {
	for i := 0; i < 10; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
}

func selectDiff() int {
	fmt.Print("Select Difficulty (Easy = 0, Normal = 1, Hard = 2) > ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "0", "1", "2":
			intInput, _ := strconv.Atoi(in)
			fmt.Print("\033[1A")
			fmt.Print("\033[K")

			switch intInput {
			case 0:
				fmt.Print("Difficulty: Easy\n")
			case 1:
				fmt.Print("Difficulty: Normal\n")
			case 2:
				fmt.Print("Difficulty: Hard\n")
			}

			return intInput

		default:
			fmt.Print("\033[1A")
			fmt.Print("\033[K")
			fmt.Print("Invalid input. Please input 0, 1 or 2 (Easy = 0, Normal = 1, Hard = 2) > ")
			continue
		}
	}
}

func displayTurn(turn int) {
	if turn == 0 {
		fmt.Print("You're first, enemy's second.\n")
	} else {
		fmt.Print("Enemy's first, You're second.\n")
	}

	fmt.Print("\n\n\n\n\n\n\n\n")
}

func displayCurrent() {
	fmt.Print("\033[8A")
	fmt.Print("┌───┬───┬───┐\n│   │   │   │\n├───┼───┼───┤\n│   │   │   │\n├───┼───┼───┤\n│   │   │   │\n└───┴───┴───┘\n")

	fmt.Print("\033[6A")
	for i := 0; i < 3; i++ {
		fmt.Print("\033[2C")

		for j := 0; j < 3; j++ {
			displayCell(state[i*3+j])
			fmt.Print("\033[3C")
		}

		fmt.Print("\n\n")
	}

	fmt.Print("\033[0m")
}

func displayCell(value string) {
	switch value {
	case "○":
		fmt.Print("\033[31m")
	case "x":
		fmt.Print("\033[36m")
	default:
		fmt.Print("\033[0m")
	}

	fmt.Print(value)
}

func playerInput() {
	fmt.Print("\033[K")
	fmt.Print("Which cell do you wanna place \"○\"? (Type number) > ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()

		ifUsed := false

		switch in {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			intInput, _ := strconv.Atoi(in)
			if state[intInput-1] == in {
				state[intInput-1] = "○"
				return
			}
			ifUsed = true
			fallthrough

		default:
			fmt.Print("\033[1A")
			fmt.Print("\033[K")
			if ifUsed {
				fmt.Print("That cell is already filled. Please input another number > ")
			} else {
				fmt.Print("Invalid input. Please input valid number > ")
			}
			continue
		}
	}
}

func ifRestart() bool {
	fmt.Print("\033[K")
	fmt.Print("Do you wanna play once more? (y/n) > ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "Y", "y":
			return true
		case "N", "n":
			return false
		default:
			fmt.Print("\033[1A")
			fmt.Print("\033[K")
			fmt.Print("Invalid input. Please input \"y\" or \"n\" > ")
			continue
		}
	}
}

func displayRes(res string) {
	fmt.Print("\033[K")
	switch res {
	case "○":
		fmt.Print("You won! Congrats! ")
	case "x":
		fmt.Print("You lose... Waiting for the next challenge! ")
	case "":
		fmt.Print("The game ended in a draw. Nice fight! ")
	}
}
