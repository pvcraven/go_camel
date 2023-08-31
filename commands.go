package main

import (
	"fmt"
	"math/rand"
)

func checkOasis(gameDataPtr *GameData) {
	roll := rand.Intn(15)
	if roll == 0 {
		fmt.Println("You found an oasis!")
		fmt.Println("You've refilled your canteen.")
		gameDataPtr.thirstLevel = 0
		gameDataPtr.drinksInCanteen = fullCanteen
	}
}

func drinkFromCanteen(gameDataPtr *GameData) {
	if gameDataPtr.drinksInCanteen <= 0 {
		fmt.Println("Your canteen is empty!")
	} else {
		fmt.Println("Awesome! You are no longer thirsty.")
		gameDataPtr.drinksInCanteen--
		gameDataPtr.thirstLevel = 0
	}
}

func travelModerateSpeed(gameDataPtr *GameData) {
	fmt.Println("Ahead moderate speed.")
	milesTraveled := rand.Intn(maxModerateMilesTraveled-minModerateMilesTraveled) + minModerateMilesTraveled
	enemyMilesTraveled := rand.Intn(maxEnemyMilesTraveled-minEnemyMilesTraveled) + minEnemyMilesTraveled
	gameDataPtr.enemyDistance += enemyMilesTraveled
	fmt.Println("You traveled", milesTraveled, "miles.")
	gameDataPtr.playerDistance += milesTraveled
	fmt.Println("You traveled a total of", gameDataPtr.playerDistance, "miles.")
	gameDataPtr.thirstLevel++
	gameDataPtr.camelFatigue++
	checkOasis(gameDataPtr)
	gameCheck(gameDataPtr)
}

func travelFullSpeed(gameDataPtr *GameData) {
	fmt.Println("Full speed ahead!")
	milesTraveled := rand.Intn(maxFullMilesTraveled-minFullMilesTraveled) + minFullMilesTraveled
	enemyMilesTraveled := rand.Intn(maxEnemyMilesTraveled-minEnemyMilesTraveled) + minEnemyMilesTraveled
	gameDataPtr.enemyDistance += enemyMilesTraveled
	fmt.Println("You traveled", milesTraveled, "miles.")
	gameDataPtr.playerDistance += milesTraveled
	gameDataPtr.thirstLevel++
	gameDataPtr.camelFatigue += rand.Intn(3) + 1

	checkOasis(gameDataPtr)
	gameCheck(gameDataPtr)
}

func stopAndRest(gameDataPtr *GameData) {
	fmt.Println("You stop and rest. Your camel thanks you.")
	enemyMilesTraveled := rand.Intn(maxEnemyMilesTraveled-minEnemyMilesTraveled) + minEnemyMilesTraveled
	gameDataPtr.enemyDistance += enemyMilesTraveled
	gameDataPtr.camelFatigue = 0
}

func quit(gameDataPtr *GameData) {
	gameDataPtr.gameOver = true
	fmt.Println("Quitting.")
}

func status(gameDataPtr *GameData) {
	fmt.Println("Status:")
	fmt.Println("  You've traveled a total of", gameDataPtr.playerDistance, "miles.")
	distanceFromPlayer := gameDataPtr.playerDistance - gameDataPtr.enemyDistance
	fmt.Println("  The natives are", distanceFromPlayer, "miles back.")
	fmt.Println("  You have", gameDataPtr.drinksInCanteen, "drinks in your canteen.")
	if gameDataPtr.camelFatigue >= 4 {
		fmt.Println("  Your camel is tired.")
	} else {
		fmt.Println("  Your camel is fine.")
	}
	if gameDataPtr.thirstLevel >= 3 {
		fmt.Println("  You are thirsty.")
	} else {
		fmt.Println("  You are not thirsty.")
	}
}
