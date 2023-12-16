package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
)

type direction int

const (
	up   direction = 0
	down   direction = 1
	left    direction = 2
	right    direction = 3
)

type beam struct {
	x, y int
	d direction
}

func main() {
	file, err := os.Open("day16input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	grid := make([][]string, 0)
	for _, line := range lines {
		row := make([]string, 0)
		for _, v := range line {
			row = append(row, string(v))
		}
		grid = append(grid, row)
	}
	// part1 := findEnergizedTiles(grid, beam{x:0, y:0, d:right})
	// fmt.Println(part1)
	activations := make([]int, 0)
	for i := range grid {
		activations = append(activations, findEnergizedTiles(grid, beam{x:i, y:0, d:right}), findEnergizedTiles(grid, beam{x:i, y:len(grid[0])-1, d:left}))
	}
	for i := range grid[0] {
		activations = append(activations, findEnergizedTiles(grid, beam{x:0, y:i, d:down}), findEnergizedTiles(grid, beam{x:len(grid)-1, y:i, d:up}))
	}
	fmt.Println(max(activations))
}

func findEnergizedTiles(grid [][]string, start beam) int {
	counts := make([][]int, 0)
	for _, line := range grid {
		countRow := make([]int, 0)
		for range line {
			countRow = append(countRow, 0)
		}
		counts = append(counts, countRow)
	}
	count := 0
	bfs := []beam{start}
	for len(bfs) != 0 {
		curr := bfs[0]
		bfs = bfs[1:]
		if curr.x >= len(grid) || curr.x < 0 || curr.y >= len(grid[0]) || curr.y < 0 {
			continue
		}
		counts[curr.x][curr.y]++
		if grid[curr.x][curr.y] == "." {
			if curr.d == up {
				bfs = append(bfs, beam{x:curr.x-1, y:curr.y, d:curr.d})
			} else if curr.d == down {
				bfs = append(bfs, beam{x:curr.x+1, y:curr.y, d:curr.d})
			} else if curr.d == left {
				bfs = append(bfs, beam{x:curr.x, y:curr.y-1, d:curr.d})
			} else if curr.d == right {
				bfs = append(bfs, beam{x:curr.x, y:curr.y+1, d:curr.d})
			}
		} else if grid[curr.x][curr.y] == "/" {
			if curr.d == up {
				bfs = append(bfs, beam{x:curr.x, y:curr.y+1, d:right})
			} else if curr.d == down {
				bfs = append(bfs, beam{x:curr.x, y:curr.y-1, d:left})
			} else if curr.d == left {
				bfs = append(bfs, beam{x:curr.x+1, y:curr.y, d:down})
			} else if curr.d == right {
				bfs = append(bfs, beam{x:curr.x-1, y:curr.y, d:up})
			}
		} else if grid[curr.x][curr.y] == "\\" {
			if curr.d == up {
				bfs = append(bfs, beam{x:curr.x, y:curr.y-1, d:left})
			} else if curr.d == down {
				bfs = append(bfs, beam{x:curr.x, y:curr.y+1, d:right})
			} else if curr.d == left {
				bfs = append(bfs, beam{x:curr.x-1, y:curr.y, d:up})
			} else if curr.d == right {
				bfs = append(bfs, beam{x:curr.x+1, y:curr.y, d:down})
			}
		} else if grid[curr.x][curr.y] == "-" {
			if curr.d == up || curr.d == down {
				bfs = append(bfs, beam{x:curr.x, y:curr.y-1, d:left}, beam{x:curr.x, y:curr.y+1, d:right})
			} else if curr.d == left {
				bfs = append(bfs, beam{x:curr.x, y:curr.y-1, d:curr.d})
			} else if curr.d == right {
				bfs = append(bfs, beam{x:curr.x, y:curr.y+1, d:curr.d})
			}
		} else if grid[curr.x][curr.y] == "|" {
			if curr.d == left || curr.d == right {
				bfs = append(bfs, beam{x:curr.x-1, y:curr.y, d:up}, beam{x:curr.x+1, y:curr.y, d:down})
			} else if curr.d == up {
				bfs = append(bfs, beam{x:curr.x-1, y:curr.y, d:curr.d})
			} else if curr.d == down {
				bfs = append(bfs, beam{x:curr.x+1, y:curr.y, d:curr.d})
			}
		}
		count++
		if count > (len(grid)*10)*(len(grid[0])*10) {
			break
		}
	}
	ans := 0
	for _, row := range counts {
		for _, v := range row {
			if v != 0 {
				ans++
			}
		}
	}
	return ans
}

func max(nums []int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}