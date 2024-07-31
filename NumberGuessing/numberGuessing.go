package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num1 int
	var num2 int
	fmt.Println("What range would you like to guess a number in?")
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	random := rand.Intn(num2-num1) + num1

	fmt.Printf("Guess a number between %d and %d.\n", num1, num2)
	var guess int
	var count int = 0
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		count++

		if guess < random {
			fmt.Println("Enter a larger number.")
		} else if guess > random {
			fmt.Println("Enter a smaller number.")
		} else {
			fmt.Printf("Congratulations! You found the number in %d steps.", count)
			break
		}
	}
}
