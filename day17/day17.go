package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
	"math"
)

type direction int

const (
	up    direction = 0
	down  direction = 1
	left  direction = 2
	right direction = 3
)

type point struct {
	x, y int
}

type result struct {
	dist, steps int
	dir direction
}

func main() {
	file, err := os.Open("day17testinput.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	grid := make([][]int, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, c := range line {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println("item in grid is not a number")
			}
			row = append(row, int(v))
		}
		grid = append(grid, row)
	}
	part1(grid)
}

func part1(grid [][]int) {
	values := make([][]int, 0)
	for _, row := range grid {
		r := make([]int, 0)
		for range row {
			r = append(r, 0)
		}
		values = append(values, r)
	}
	numNodes := len(grid)*len(grid[0])
	visited := make(map[point]result)
	visited[point{x:0, y:0}] = result{dist:0, steps:0, dir:right}
	for len(visited) != numNodes {
		fmt.Println(visited)
		minNode := point{}
		minNeighbour := point{}
		minVal := math.MaxInt
		for k, v := range visited {
			neighbours := getNeighbours(grid, k)
			for _, n := range neighbours {
				if _, ok := visited[n]; ok {
					continue
				}
				if getDirection(k, n) == v.dir && v.steps >= 3 {
					continue
				}
				if grid[n.x][n.y] + v.dist < minVal {
					minVal = grid[n.x][n.y] + v.dist
					minNode = k
					minNeighbour = n
				}
			}
		}
		minDir := getDirection(minNode, minNeighbour)
		value := result{}
		if minDir == visited[minNode].dir {
			value = result{dist:minVal, dir:minDir, steps:visited[minNode].steps+1}
		} else {
			value = result{dist:minVal, dir:minDir, steps:1}
		}
		visited[minNeighbour] = value
		fmt.Println(visited)
	}
	fmt.Println(visited[point{x:len(grid)-1, y:len(grid[0])-1}].dist)
}

func getNeighbours(grid [][]int, node point) []point {
	neighbours := make([]point, 0)
	x := node.x
	y := node.y
	if x > 0 {
		neighbours = append(neighbours, point{x-1, y})
	}
	if x < len(grid)-1 {
		neighbours = append(neighbours, point{x+1, y})
	}
	if y > 0 {
		neighbours = append(neighbours, point{x, y-1})
	}
	if y < len(grid[0])-1 {
		neighbours = append(neighbours, point{x, y+1})
	}
	return neighbours
}

// Direction of a relative to b. Only works for adjacent points
func getDirection(a point, b point) direction {
	if a.x > b.x {
		return up
	} else if a.x < b.x {
		return down
	} else if a.y > b.y {
		return left
	}
	return right
}