package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type game struct {
	ID     int
	reds   int
	blues  int
	greens int
}

func (g game) IsPossibleGame(redLimit int, greenLimit int, blueLimit int) bool {
	if g.reds <= redLimit && g.greens <= greenLimit && g.blues <= blueLimit {
		return true
	}
	return false
}

func (g game) GetPower() int {
	return g.reds * g.greens * g.blues
}

func main() {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	file, err := os.Open("day2input.txt")
	if err != nil {
		panic(err)
	}

	// Read the file contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Close the file
	defer file.Close()

	var games []game
	for _, v := range strings.Split(string(contents), "\n") {
		gameObj := parseGame(v)
		games = append(games, gameObj)
	}

	total := 0
	totalPower := 0
	for _, gameObj := range games {
		if gameObj.IsPossibleGame(redLimit, greenLimit, blueLimit) {
			total += gameObj.ID
		}
		totalPower += gameObj.GetPower()
	}
	fmt.Println(total)
	fmt.Println(totalPower)
}

func parseGame(gameString string) game {
	gameIDAndValues := strings.Split(gameString, ":")
	gameID, err := strconv.Atoi(strings.Split(gameIDAndValues[0], " ")[1])
	if err != nil {
		fmt.Printf("game id is not an integer: %d\n", gameID)
	}
	rounds := strings.Split(gameIDAndValues[1], ";")
	maxRed := 0
	maxBlue := 0
	maxGreen := 0

	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			data := strings.Split(cube, " ")
			numCubes, err := strconv.Atoi(string(data[1]))
			if err != nil {
				fmt.Println("number of cubes is not a number")
			}
			switch data[2] {
			case "red":
				maxRed = max(maxRed, numCubes)
			case "blue":
				maxBlue = max(maxBlue, numCubes)
			case "green":
				maxGreen = max(maxGreen, numCubes)
			}
		}

	}

	return game{ID: gameID, reds: maxRed, blues: maxBlue, greens: maxGreen}
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
