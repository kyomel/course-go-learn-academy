package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	// print to screen
	// keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 0
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			// changed here to send into back when done
			g.DisplayChan <- ""
		}
	}
}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {
	// print out some instructions
	// fmt.Println("Rock, Paper & Scissors")
	// fmt.Println("-----------------------")
	// fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	// fmt.Println()

	// changed here
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
-----------------------
Game is played for three rounds, and best of three wins the game. Good luck!
	`)
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	// fmt.Println()
	// fmt.Println("Round", g.Round.RoundNumber)
	// fmt.Println("------")
	// changed here
	g.DisplayChan <- fmt.Sprintf(`
Round %d
---------
	`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Printf("Please enter rock, paper, or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	// print out player choice
	// changed here
	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf(`Player choose %s`, strings.ToUpper(playerChoice))
	<-g.DisplayChan

	// If matching boolean first that will be ignore everything else with same value
	switch computerValue {
	case ROCK:
		// changed here
		g.DisplayChan <- "Computer choose ROCK"
		<-g.DisplayChan
		break
	case PAPER:
		g.DisplayChan <- "Computer choose PAPER"
		<-g.DisplayChan
		break
	case SCISSORS:
		g.DisplayChan <- "Computer choose SCISSORS"
		<-g.DisplayChan
		break
	default:
	}

	// changed here
	g.DisplayChan <- ""
	<-g.DisplayChan

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		// changed here
		<-g.DisplayChan
		return false
	} else {
		// refactor the logic to keep track of score and print
		// messages to two new functions, computer wins and player wins
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice!"
			// changed here
			<-g.DisplayChan
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	// changed here
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	// changed here
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	// fmt.Println("Final score")
	// fmt.Println("--------------")
	// fmt.Printf("Player: %d/3, computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	// fmt.Println()
	// changed here
	g.DisplayChan <- fmt.Sprintf(`
Final Score
------------
Player: %d/3, Computer %d/3
	`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan

	if g.Round.PlayerScore > g.Round.ComputerScore {
		// fmt.Println("Player wins the game!")
		// changed here
		g.DisplayChan <- "Player wins game!"
		<-g.DisplayChan
	} else {
		// fmt.Println("Computer wins game")
		// changed here
		g.DisplayChan <- "Computer wins game!"
		<-g.DisplayChan
	}

	// changed here
	g.DisplayChan <- ""
	<-g.DisplayChan
}
