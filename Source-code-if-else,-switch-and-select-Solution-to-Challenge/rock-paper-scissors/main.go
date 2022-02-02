package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1
	// *** added two variables to keep track of score
	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	// *** print out some instructions
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()

	// *** added the for loop
	for i := 1; i <= 3; i++ {
		// *** print out round information
		fmt.Println()
		fmt.Println("Round", i)
		fmt.Println("-------")

		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		// moved computer choice to for loop, so it's reset each time through
		computerValue := rand.Intn(3)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			playerValue = -1
		}

		// *** print out player choice
		fmt.Println()
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
			break
		case PAPER:
			fmt.Println("Computer chose PAPER")
			break
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
			break
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
			// *** decrement i by 1, since we're repeating the round
			i--
		} else {
			// *** refactor the logic to keep track of score and print
			// messages to two new functions, computerWins and playerWins
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			default:
				// *** decrement i by 1, since we're repeating the round
				fmt.Println("Invalid choice!")
				i--
			}
		}
	}

	fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
	fmt.Println()
	if playerScore > computerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}
}

func computerWins(score int) int {
	fmt.Println("Computer wins!")
	return score + 1
}

func playerWins(score int) int {
	fmt.Println("Player wins!")
	return score + 1
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
