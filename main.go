package main

import (
	"fmt"
	"math/rand"
	"time"
)

var userWins = 0
var computerWins = 0
var choices = []string{"КАМЕНЬ", "НОЖНИЦЫ", "БУМАГА"}

func main() {
	fmt.Println("Камень, Ножницы, Бумага!")

	for {
		fmt.Println("Сделайте ваш выбор:")
		fmt.Println("1-->КАМЕНЬ")
		fmt.Println("2-->НОЖНИЦЫ")
		fmt.Println("3-->БУМАГА")

		var userInput int
		fmt.Scan(&userInput)

		userChoice := choices[userInput-1]
		computerChoice := generateComputerChoice()

		findWinner(userChoice, computerChoice)

		var input string
		fmt.Println("Еще раз (д/н)?:")
		fmt.Scan(&input)

		if input == "н" {
			break
		}
	}

	fmt.Println("Игра окончена")

	fmt.Printf("Вы выиграли %d раз", userWins)
	fmt.Println()
	fmt.Printf("Вы проиграли %d раз", computerWins)
}
func generateComputerChoice() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	index := random.Intn(3)

	return choices[index]
}

func findWinner(userChoice, computerChoice string) {
	if userChoice == computerChoice {
		fmt.Println("Ничья")
	} else if userChoice == "КАМЕНЬ" && computerChoice == "НОЖНИЦЫ" {
		userWins++
		fmt.Println("Вы выиграли!")
	} else if userChoice == "НОЖНИЦЫ" && computerChoice == "БУМАГА" {
		userWins++
		fmt.Println("Вы выиграли!")
	} else if userChoice == "БУМАГА" && computerChoice == "КАМЕНЬ" {
		userWins++
		fmt.Println("Вы выиграли!")
	} else {
		computerWins++
		fmt.Println("Вы проиграли!")
	}
}
