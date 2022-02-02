package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("-> ")
	// 	userInput, _ := reader.ReadString('\n')
	// 	userInput = strings.Replace(userInput, "\n", "", -1)
	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	char := ' '
	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(string(char))
		_, ok := coffees[i]
		if ok {
			fmt.Println(fmt.Sprintf("You choose %s", coffees[i]))
		}
		// part way
		// if _, ok := coffees[i]; ok {
		// 	fmt.Println(fmt.Sprintf("You choose %s", coffees[i]))
		// }
	}

	// other way
	// for ok := true; ok; ok =char != 'q' && char != 'Q' {
	// 	char, _, err = keyboard.GetSingleKey()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	i, _ := strconv.Atoi(string(char))
	// 	_, ok := coffees[i]
	// 	if ok {
	// 		fmt.Println(fmt.Sprintf("You choose %s", coffees[i]))
	// 	}
	// }
	fmt.Println("Program exiting...")
}

// better for Go
// A
// for i := 0 <= 100; i = i + 7 {
// 	fmt.Println(i)
// }
// B
// for i := 0; i <= 100; i++ {
// 	if i % 7 == 0 {
// 		fmt.Println(i)
// 	}
// }
