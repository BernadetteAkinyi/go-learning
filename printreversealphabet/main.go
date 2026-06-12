package main

import "fmt"

func main() {
	var choice int

	fmt.Println("===GO TOOLKIT ===")
	fmt.Println("1. Print reverse alphabet")
	fmt.Println("2. Print alphabet")
	fmt.Println("3. Print numbers 1 to 10 ")
	fmt.Println("4. Exit")

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	fmt.Println("You selected:", choice)

	if choice == 1 {
		for c := 'z'; c >= 'a'; c-- {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	} else if choice == 2 {
		for c := 'a'; c <= 'z'; c++ {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	} else if choice == 3 {
		for b := 1; b <= 10; b++ {
			fmt.Println(b)
		}
		fmt.Println()

	} else if choice == 4 {
		fmt.Println("Goodbye")
	} else {
		fmt.Println("Invalid choice")
	}

}
