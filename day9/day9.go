package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	sequences := getSequences()
	total := 0
	for _, sequence := range sequences {
		total += computeNextValue(sequence)
	}
	fmt.Println(total)
}

func part2() {
	sequences := getSequences()
	total := 0
	for _, sequence := range sequences {
		total += computePreviousValue(sequence)
	}
	fmt.Println(total)
}

func getSequences() [][]int {
	file, err := os.Open("day9input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	sequences := make([][]int, 0)
	for _, line := range lines {
		nums := strings.Split(line, " ")
		sequence := make([]int, 0)
		for _, num := range nums {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("number is not a number")
			}
			sequence = append(sequence, numInt)
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func computeNextValue(sequence []int) int {
	if sum(sequence) == 0 {
		return 0
	}
	diffs := make([]int, 0)
	for i := 0; i < len(sequence)-1; i++ {
		diffs = append(diffs, sequence[i+1]-sequence[i])
	}
	nextVal := computeNextValue(diffs)
	return sequence[len(sequence)-1] + nextVal
}

func computePreviousValue(sequence []int) int {
	if sum(sequence) == 0 {
		return 0
	}
	diffs := make([]int, 0)
	for i := 0; i < len(sequence)-1; i++ {
		diffs = append(diffs, sequence[i+1]-sequence[i])
	}
	prevVal := computePreviousValue(diffs)
	return sequence[0] - prevVal
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
