package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type direction string

const (
	north   direction = "coming from north"
	south   direction = "coming from south"
	east    direction = "coming from east"
	west    direction = "coming from west"
	neutral direction = "coming from start"
)

type point struct {
	x, y int
}

func main() {
	file, err := os.Open("day10input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	startPoint := point{}
	grid := make([][]string, 0)
	cleanGrid := make([][]string, 0)
	for i, line := range lines {
		chars := make([]string, 0)
		cleanChars := make([]string, 0)
		for j, char := range line {
			chars = append(chars, string(char))
			cleanChars = append(cleanChars, ".")
			if char == 'S' {
				startPoint = point{x: i, y: j}
			}
		}
		grid = append(grid, chars)
		cleanGrid = append(cleanGrid, cleanChars)
	}
	cleanGrid[startPoint.x][startPoint.y] = "S"
	currentPoint, currentDirection := getNextMove(grid, startPoint, neutral)
	distance := 1
	for currentDirection != neutral {
		cleanGrid[currentPoint.x][currentPoint.y] = grid[currentPoint.x][currentPoint.y]
		currentPoint, currentDirection = getNextMove(grid, currentPoint, currentDirection)
		distance++
	}
	cleanGrid = replaceStart(cleanGrid, startPoint)
	for _, line := range cleanGrid {
		fmt.Println(line)
	}
	fmt.Println(distance / 2)
	fmt.Println(getAreaOfNest(cleanGrid))
}

func getAreaOfNest(grid [][]string) int {
	insideNest := false
	lastBend := ""
	area := 0
	for _, row := range grid {
		for _, pipe := range row {
			if pipe == "." && insideNest {
				area++
			} else if pipe == "|" {
				insideNest = !insideNest
			} else if lastBend == "" && pipeIsBend(pipe) {
				lastBend = pipe
			} else if lastBend != "" && pipeIsBend(pipe) && ((lastBend == "L" && pipe == "7") || (lastBend == "F" && pipe == "J")) {
				insideNest = !insideNest
				lastBend = ""
			} else if lastBend != "" && pipeIsBend(pipe) && ((lastBend == "L" && pipe == "J") || (lastBend == "F" && pipe == "7")) {
				lastBend = ""
			}
		}
	}
	return area
}

func pipeIsBend(pipe string) bool {
	return pipe == "L" || pipe == "J" || pipe == "F" || pipe == "7"
}

func getNextMove(grid [][]string, current point, incoming direction) (point, direction) {
	x := current.x
	y := current.y
	if grid[x][y] == "S" {
		if y > 0 && canGoEast(grid[x][y-1]) {
			nextPoint := point{x: x, y: y - 1}
			return nextPoint, east
		} else if y < len(grid[0])-1 && canGoWest(grid[x][y+1]) {
			nextPoint := point{x: x, y: y + 1}
			return nextPoint, west
		} else if x > 0 && canGoSouth(grid[x-1][y]) {
			nextPoint := point{x: x - 1, y: y}
			return nextPoint, south
		} else if x < len(grid)-1 && canGoNorth(grid[x+1][y]) {
			nextPoint := point{x: x + 1, y: y}
			return nextPoint, south
		}
	} else if incoming != north && x > 0 && canGoNorth(grid[x][y]) && canGoSouth(grid[x-1][y]) {
		nextPoint := point{x: x - 1, y: y}
		return nextPoint, south
	} else if incoming != south && x < len(grid)-1 && canGoSouth(grid[x][y]) && canGoNorth(grid[x+1][y]) {
		nextPoint := point{x: x + 1, y: y}
		return nextPoint, north
	} else if incoming != west && y > 0 && canGoWest(grid[x][y]) && canGoEast(grid[x][y-1]) {
		nextPoint := point{x: x, y: y - 1}
		return nextPoint, east
	} else if incoming != east && y < len(grid[0])-1 && canGoEast(grid[x][y]) && canGoWest(grid[x][y+1]) {
		nextPoint := point{x: x, y: y + 1}
		return nextPoint, west
	}
	return point{}, neutral
}

func canGoNorth(current string) bool {
	return current == "|" || current == "J" || current == "L"
}

func canGoSouth(current string) bool {
	return current == "|" || current == "F" || current == "7"
}

func canGoEast(current string) bool {
	return current == "-" || current == "L" || current == "F"
}

func canGoWest(current string) bool {
	return current == "-" || current == "J" || current == "7"
}

func replaceStart(grid [][]string, start point) [][]string {
	x := start.x
	y := start.y
	southOpen := false
	northOpen := false
	eastOpen := false
	westOpen := false
	if x > 0 && canGoSouth(grid[x-1][y]) {
		northOpen = true
	}
	if x < len(grid)-1 && canGoNorth(grid[x+1][y]) {
		southOpen = true
	}
	if y > 0 && canGoEast(grid[x][y-1]) {
		westOpen = true
	}
	if y < len(grid[0])-1 && canGoWest(grid[x][y+1]) {
		eastOpen = true
	}
	if northOpen && southOpen {
		grid[x][y] = "|"
	} else if northOpen && eastOpen {
		grid[x][y] = "L"
	} else if northOpen && westOpen {
		grid[x][y] = "J"
	} else if southOpen && westOpen {
		grid[x][y] = "7"
	} else if southOpen && eastOpen {
		grid[x][y] = "F"
	} else if westOpen && eastOpen {
		grid[x][y] = "-"
	} else {
		panic("Start doesn't connect to anything")
	}
	return grid
}
