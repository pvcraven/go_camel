package main

// Holds the data specific to any instance of the game.
type GameData struct {
	playerDistance  int
	enemyDistance   int
	camelFatigue    int
	drinksInCanteen int
	thirstLevel     int
	gameOver        bool
}

func createGame() GameData {
	var gameData GameData
	gameData.playerDistance = 0
	gameData.enemyDistance = enemyStartingLocation
	gameData.gameOver = false
	gameData.drinksInCanteen = fullCanteen
	return gameData
}
