package main

import (
	"bufio"
	"fmt"
	"gobasic/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		fmt.Println(doctor.Response(userInput))

		if userInput == "quit" {
			break
		}
	}
}
