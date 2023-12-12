package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"math"
)

type galaxy struct {
	id, x, y int
}

func main() {
	file, err := os.Open("day11input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	universe := expandUniverse(lines)
	galaxyList := make([]galaxy, 0)
	galaxyNum := 1
	for i, line := range universe {
		for j, char := range line {
			if char == "#" {
				newGalaxy := galaxy{x:i, y:j, id:galaxyNum}
				galaxyList = append(galaxyList, newGalaxy)
				galaxyNum++
			}
		}
	}
	total := 0.0
	for i := 0; i < len(galaxyList)-1; i++ {
		for j := i+1; j < len(galaxyList); j++ {
			total += getDistance(galaxyList[i], galaxyList[j])
		}
	}
	fmt.Println(int(total))
}

func part2(lines []string) {
	expandedRows := make([]int, 0)
	for i, line := range lines {
		if !strings.Contains(line, "#") {
			expandedRows = append(expandedRows, i)
		}
	}
	galaxyList := make([]galaxy, 0)
	galaxyNum := 1
	universeColMap := make(map[int]bool, 0)
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				universeColMap[j] = true
				newGalaxy := galaxy{x:i, y:j, id:galaxyNum}
				galaxyList = append(galaxyList, newGalaxy)
				galaxyNum++
			}
		}
	}
	expandedCols := make([]int, 0)
	for i := range lines[0] {
		if _, ok := universeColMap[i]; !ok {
			expandedCols = append(expandedCols, i)
		}
	}
	total := 0
	for i := 0; i < len(galaxyList)-1; i++ {
		for j := i+1; j < len(galaxyList); j++ {
			total += getExpansionsBetweenPoints(galaxyList[i], galaxyList[j], expandedRows, expandedCols)
		}
	}
	fmt.Println(total)
}

func getExpansionsBetweenPoints(galaxy1, galaxy2 galaxy, expandedRows, expandedCols []int) int {
	distance := int(getDistance(galaxy1, galaxy2))
	bigX := galaxy2.x
	bigY := galaxy2.y
	smallX := galaxy1.x
	smallY := galaxy1.y
	if galaxy1.x > galaxy2.x {
		bigX = galaxy1.x
		smallX = galaxy2.x
	}
	if galaxy1.y > galaxy2.y {
		bigY = galaxy1.y
		smallY = galaxy2.y
	}
	for _, i := range expandedRows {
		if smallX < i && i < bigX {
			distance += 999999
		}
	}
	for _, i := range expandedCols {
		if smallY < i && i < bigY {
			distance += 999999
		}
	}
	return distance
}

func getDistance(galaxy1 galaxy, galaxy2 galaxy) float64 {
	return math.Abs(float64(galaxy2.y - galaxy1.y)) + math.Abs(float64(galaxy2.x - galaxy1.x))
}

func expandUniverse(lines []string) [][]string {
	tallUniverse := make([]string, 0)
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			tallUniverse = append(tallUniverse, line)
		}
		tallUniverse = append(tallUniverse, line)
	}
	universeCols := make(map[int]bool, 0)
	for _, line := range lines {
		for i, char := range line {
			if char == '#' {
				universeCols[i] = true
			}
		}
	}
	fullUniverse := make([][]string, 0)
	for _, line := range tallUniverse {
		row := make([]string, 0)
		for j, char := range line {
			if _, ok := universeCols[j]; !ok {
				row = append(row, string(char))
			}
			row = append(row, string(char))
		}
		fullUniverse = append(fullUniverse, row)
	}
	return fullUniverse
}

func expandUniverseALot(lines []string) [][]string {
	tallUniverse := make([]string, 0)
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			for i := 0; i < 999999; i++ {
				tallUniverse = append(tallUniverse, line)
			}
		}
		tallUniverse = append(tallUniverse, line)
	}
	universeCols := make(map[int]bool, 0)
	for _, line := range lines {
		for i, char := range line {
			if char == '#' {
				universeCols[i] = true
			}
		}
	}
	fullUniverse := make([][]string, 0)
	for _, line := range tallUniverse {
		row := make([]string, 0)
		for j, char := range line {
			if _, ok := universeCols[j]; !ok {
				for i := 0; i < 999999; i++ {
					row = append(row, string(char))
				}
			}
			row = append(row, string(char))
		}
		fullUniverse = append(fullUniverse, row)
	}
	return fullUniverse
}