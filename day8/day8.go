package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {
	part2()
}

func part2() {
	graph, instructions := createGraphAndInstructions()
	// steps := 0
	startingNodes := make([]string, 0)
	for key := range graph {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}
	times := make([]int, 0)
	for _, startingNode := range startingNodes {
		times = append(times, traversePath(graph, instructions, startingNode))
	}
	fmt.Println(times)
	runningLcm := times[0]
	for i := 1; i < len(times); i++ {
		runningLcm = lcm(runningLcm, times[i])
	}
	fmt.Println(runningLcm)

}

func traversePath(graph map[string]node, instructions string, start string) int {
	i := 0
	currentPosition := start
	for true {
		if currentPosition[2] == 'Z' {
			break
		}
		turn := instructions[i%len(instructions)]
		if turn == 'L' {
			currentPosition = graph[currentPosition].left
		} else if turn == 'R' {
			currentPosition = graph[currentPosition].right
		}
		i++
	}
	return i
}

func part1() {
	graph, instructions := createGraphAndInstructions()
	i := 0
	currentPosition := "AAA"
	for true {
		if currentPosition == "ZZZ" {
			break
		}
		turn := instructions[i%len(instructions)]
		if turn == 'L' {
			currentPosition = graph[currentPosition].left
		} else if turn == 'R' {
			currentPosition = graph[currentPosition].right
		}
		i++
	}
	fmt.Println(i)
}

func createGraphAndInstructions() (map[string]node, string) {
	file, err := os.Open("day8input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	instructions := lines[0]
	graph := make(map[string]node)
	for _, line := range lines[2:] {
		parts := strings.Split(line, "=")
		source := parts[0][:3]
		destination := strings.Split(parts[1], ",")
		left := destination[0][2:]
		right := destination[1][1:4]
		graph[source] = node{left: left, right: right}
	}
	return graph, instructions
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
