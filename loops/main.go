package main

import "fmt"

func main() {
	// three parts loops
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("i is", i)
	// }

	// for i := 10; i >= 0; i-- {
	// 	fmt.Println("i is", i)
	// }

	// execute a loop while i > 100
	// rand.Seed(time.Now().UnixNano())
	// i := 1000
	// for i > 100 {
	// 	//  get a random number between 1 and 1001
	// 	i = rand.Intn(1000) + 1
	// 	fmt.Println("i is", i)
	// 	if i > 100 {
	// 		fmt.Println("i is", i, "so loop keeps going")
	// 	}
	// }
	// fmt.Println("Got", i, "and broke out of loop")

	// read input from user 5 times and write it to a log
	// infinite loops
	// reader := bufio.NewReader(os.Stdin)
	// ch := make(chan string)

	// go mylogger.ListenForLog(ch)

	// fmt.Println("Enter something")

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("-> ")
	// 	input, _ := reader.ReadString('\n')
	// 	ch <- input
	// 	time.Sleep(time.Second)
	// }

	for i := 1; i <= 10; i++ {
		fmt.Print("i is", i)
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j)
		}
		fmt.Println()
	}
}
