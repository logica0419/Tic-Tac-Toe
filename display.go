package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func displayInit() {
	fmt.Print("Lets Play Tic-Tac-Toe!\nPreparing...")
	for i := 0; i < 10; i++ {
		fmt.Print("...")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("\n")
	for i := 0; i < 2; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
	fmt.Print("\n\n\n\n\n\n\n\n")
}

func displayRestart() {
	for i := 0; i < 8; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
	fmt.Print("Restarting Tic-Tac-Toe!\nPreparing...")
	for i := 0; i < 10; i++ {
		fmt.Print("...")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("\n")
	for i := 0; i < 2; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
	fmt.Print("\n\n\n\n\n\n\n\n")
}

func displayCurrent() {
	fmt.Print("\033[8A")
	field := fmt.Sprintf("┌───┬───┬───┐\n│ %s │ %s │ %s │\n├───┼───┼───┤\n│ %s │ %s │ %s │\n├───┼───┼───┤\n│ %s │ %s │ %s │\n└───┴───┴───┘\n",
		state[0], state[1], state[2], state[3], state[4], state[5], state[6], state[7], state[8])
	fmt.Print(field)
}

func playerInput() {
	fmt.Print("\033[K")
	fmt.Print("Where do you wanna place \"○\"? > ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()

		ifUsed := false

		switch in {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			inputNum, _ := strconv.Atoi(in)
			if state[inputNum-1] == in {
				state[inputNum-1] = "○"
				return
			}
			ifUsed = true
			fallthrough

		default:
			fmt.Print("\033[1A")
			fmt.Print("\033[K")
			if ifUsed {
				fmt.Print("That place is already filled. Please input another number > ")
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
		fmt.Print("You lose... Waiting for your next challenge! ")
	case "":
		fmt.Print("The game is draw. Nice fight! ")
	}
}
