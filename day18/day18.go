package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction string

type instruction struct {
	dir  string
	dist int
}

type point struct {
	x, y int
}

func main() {
	file, err := os.Open("day18input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	part2(lines)
}

func part1(lines []string) {
	instructions := getInstructions(lines)
	grid := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		row := make([]string, 0)
		for j := 0; j < 1000; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}
	grid[500][500] = "#"
	curr := point{x: 500, y: 500}
	for _, inst := range instructions {
		if inst.dir == "R" {
			i := curr.y + 1
			for i <= curr.y+inst.dist {
				grid[curr.x][i] = "#"
				i++
			}
			curr = point{x: curr.x, y: i - 1}
		} else if inst.dir == "L" {
			i := curr.y - 1
			for i >= curr.y-inst.dist {
				grid[curr.x][i] = "#"
				i--
			}
			curr = point{x: curr.x, y: i + 1}
		} else if inst.dir == "U" {
			i := curr.x - 1
			for i >= curr.x-inst.dist {
				grid[i][curr.y] = "#"
				i--
			}
			curr = point{x: i + 1, y: curr.y}
		} else if inst.dir == "D" {
			i := curr.x + 1
			for i <= curr.x+inst.dist {
				grid[i][curr.y] = "#"
				i++
			}
			curr = point{x: i - 1, y: curr.y}
		}
	}
	bfs(grid)
	fmt.Println(findArea(grid))
}

func part2(lines []string) {
	instructions := decodeInstructions(lines)
	points := make([]point, 0)
	points = append(points, point{x: 0, y: 0})
	currPoint := point{x: 0, y: 0}
	perimeter := 0
	for _, inst := range instructions {
		perimeter += inst.dist
		if inst.dir == "R" {
			newPoint := point{x: currPoint.x, y: currPoint.y + inst.dist}
			currPoint = newPoint
			points = append(points, newPoint)
		} else if inst.dir == "L" {
			newPoint := point{x: currPoint.x, y: currPoint.y - inst.dist}
			currPoint = newPoint
			points = append(points, newPoint)
		} else if inst.dir == "U" {
			newPoint := point{x: currPoint.x - inst.dist, y: currPoint.y}
			currPoint = newPoint
			points = append(points, newPoint)
		} else if inst.dir == "D" {
			newPoint := point{x: currPoint.x + inst.dist, y: currPoint.y}
			currPoint = newPoint
			points = append(points, newPoint)
		}
	}
	area := shoelaceFormula(points)
	ans := area + perimeter/2 + 1
	fmt.Println(ans)
}

func decodeInstructions(lines []string) []instruction {
	hexToDir := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}
	instructions := make([]instruction, 0)
	for _, line := range lines {
		v := strings.Split(line, " ")
		color := v[2][2 : len(v[2])-1]
		decimalNum, err := strconv.ParseInt(color[:len(color)-1], 16, 64)
		if err != nil {
			fmt.Println("hex is not a number")
		}
		dir := hexToDir[string(color[len(color)-1])]
		newInstruction := instruction{dir: dir, dist: int(decimalNum)}
		instructions = append(instructions, newInstruction)
	}
	return instructions
}

func getInstructions(lines []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range lines {
		v := strings.Split(line, " ")
		dir := v[0]
		dist, err := strconv.Atoi(v[1])
		if err != nil {
			fmt.Println("distance is not a number")
		}
		newInstruction := instruction{dist: dist, dir: dir}
		instructions = append(instructions, newInstruction)
	}
	return instructions
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func bfs(grid [][]string) {
	queue := []point{point{x: 0, y: 0}}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.x < 0 || curr.y < 0 || curr.x > len(grid)-1 || curr.y > len(grid[0])-1 {
			continue
		}
		if grid[curr.x][curr.y] == "#" || grid[curr.x][curr.y] == "$" {
			continue
		}
		if grid[curr.x][curr.y] == "." {
			grid[curr.x][curr.y] = "$"
		}
		queue = append(queue, point{curr.x - 1, curr.y}, point{curr.x + 1, curr.y}, point{curr.x, curr.y - 1}, point{curr.x, curr.y + 1})
	}
}

func findArea(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v == "#" || v == "." {
				count++
			}
		}
	}
	return count
}

func shoelaceFormula(points []point) int {
	numPoints := len(points)
	sum1, sum2 := 0, 0
	for i := range points[:numPoints-1] {
		sum1 += points[i].x * points[i+1].y
		sum2 += points[i].y * points[i+1].x
	}
	fmt.Println(sum1, sum2)
	area := math.Abs(float64(sum1)-float64(sum2)) / 2
	return int(area)
}
