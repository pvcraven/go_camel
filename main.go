/*
This is a runnable program that implements the Camal game in the Go language.
See: https://learn.arcade.academy/en/latest/labs/lab_04_camel/camel.html
*/

package main

import (
	"fmt"
	"strings"
)

func printWelcome() {
	fmt.Println("Welcome to Camel!")
	fmt.Println("You have stolen a camel to make your way across the great Mobi desert.")
	fmt.Println("The natives want their camel back and are chasing you down! Survive your")
	fmt.Println("desert trek and outrun the natives.")
	fmt.Println()
}

func printMenu() {
	fmt.Println("A. Drink from your canteen.")
	fmt.Println("B. Ahead moderate speed.")
	fmt.Println("C. Ahead full speed.")
	fmt.Println("D. Stop and rest.")
	fmt.Println("E. Status check.")
	fmt.Println("Q. Quit.")
}

func getUserChoice() string {

	done := false
	var userInput string
	validAnswers := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"q": true,
	}

	for !done {
		fmt.Print("Your choice? ")
		fmt.Scan(&userInput)
		userInput = strings.ToLower(userInput)
		if validAnswers[userInput] {
			done = true
		} else {
			fmt.Print("\nInvalid choice.\n\n")
			printMenu()
		}
	}
	return userInput
}

func gameCheck(gameDataPtr *GameData) {
	// Enemy distance
	distanceFromPlayer := gameDataPtr.playerDistance - gameDataPtr.enemyDistance
	if distanceFromPlayer <= 0 {
		fmt.Println("Oh no! The natives have caught up with you.")
		gameDataPtr.gameOver = true
	} else if distanceFromPlayer <= warningDistance {
		fmt.Println("The natives are getting close!")
	}
	// Thirst
	if gameDataPtr.thirstLevel > 5 {
		fmt.Println("You died of thirst.")
		gameDataPtr.gameOver = true
	} else if gameDataPtr.thirstLevel > 3 {
		fmt.Println("You are thirsty.")
	}
	// Camel
	if gameDataPtr.camelFatigue >= 6 {
		fmt.Println("Your camel is died of exhaustion.")
		fmt.Println("You died wandering the desert.")
		gameDataPtr.gameOver = true
	} else if gameDataPtr.camelFatigue >= 4 {
		fmt.Println("Your camel is tired.")
	}
	// Finish line
	if !gameDataPtr.gameOver && gameDataPtr.playerDistance > distanceToWin {
		fmt.Println("*** Congratulations ***")
		fmt.Println("You made it safely across the desert!")
		gameDataPtr.gameOver = true
	}
}

func main() {
	printWelcome()
	printMenu()

	gameData := createGame()
	gameDataPtr := &gameData

	for !gameData.gameOver {
		userInput := getUserChoice()
		fmt.Println()

		if userInput == "a" {
			drinkFromCanteen(gameDataPtr)
		} else if userInput == "b" {
			travelModerateSpeed(gameDataPtr)
		} else if userInput == "c" {
			travelFullSpeed(gameDataPtr)
		} else if userInput == "d" {
			stopAndRest(gameDataPtr)
		} else if userInput == "q" {
			quit(gameDataPtr)
		} else if userInput == "e" {
			status(gameDataPtr)
		}
	}
}
