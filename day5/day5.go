package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type numberRange struct {
	sourceRangeStart      int
	destinationRangeStart int
	rangeLength           int
}

type mapping struct {
	name   string
	ranges []numberRange
}

func main() {
	mappings, seeds := buildMappings()
	part1(mappings, seeds)
	part2(mappings, seeds)
}

func part1(mappings []mapping, seeds []int) {
	mappedSeeds := make([]int, 0)
	for _, seed := range seeds {
		mappedSeed := mapSeeds(seed, mappings)
		mappedSeeds = append(mappedSeeds, mappedSeed)
	}

	fmt.Println(min(mappedSeeds))
}

func part2(mappings []mapping, seeds []int) {
	seedRanges := make([]numberRange, 0)
	currentRange := numberRange{}
	for i, seed := range seeds {
		if i%2 == 0 {
			currentRange.sourceRangeStart = seed
		} else {
			currentRange.rangeLength = seed
			seedRanges = append(seedRanges, currentRange)
			currentRange = numberRange{}
		}
	}
	for i := 0; i < math.MaxInt; i++ {
		currentSeed := mapLocationToSeed(i, mappings)
		for _, seedRange := range seedRanges {
			if currentSeed >= seedRange.sourceRangeStart && currentSeed < seedRange.sourceRangeStart + seedRange.rangeLength {
				fmt.Println(i)
				return
			}
		}
	}

}

func buildMappings() ([]mapping, []int) {
	file, err := os.Open("day5input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	seeds := strings.Fields(lines[0])[1:]
	var filteredLines []string
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	lines = filteredLines
	var mappings []mapping
	currentMapping := mapping{}
	for _, line := range lines[1:] {
		if string(line[len(line)-1]) == ":" {
			if currentMapping.name == "" {
				currentMapping.name = line
			} else {
				mappings = append(mappings, currentMapping)
				currentMapping = mapping{name: line}
			}
		} else {
			numbersString := strings.Split(line, " ")
			var numbers []int
			for i := range numbersString {
				number, err := strconv.Atoi(string(numbersString[i]))
				if err != nil {
					fmt.Println("range value is not a number")
				}
				numbers = append(numbers, number)
			}
			currentNumberRange := numberRange{destinationRangeStart: numbers[0], sourceRangeStart: numbers[1], rangeLength: numbers[2]}
			currentMapping.ranges = append(currentMapping.ranges, currentNumberRange)
		}
	}
	mappings = append(mappings, currentMapping)

	convertedSeeds := make([]int, 0)
	for _, seed := range seeds {
		convertedSeed, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("seed is not a number")
		}
		convertedSeeds = append(convertedSeeds, convertedSeed)
	}

	return mappings, convertedSeeds
}

func mapSeeds(seed int, mappings []mapping) int {
	currentValue := seed
	for _, currentMapping := range mappings {
	rangeLoop:
		for _, currentRange := range currentMapping.ranges {
			sourceFloor := currentRange.sourceRangeStart
			sourceCeiling := currentRange.sourceRangeStart + currentRange.rangeLength
			if currentValue >= sourceFloor && currentValue < sourceCeiling {
				diff := currentValue - sourceFloor
				currentValue = currentRange.destinationRangeStart + diff
				break rangeLoop
			}
		}
	}
	return currentValue
}

func mapLocationToSeed(location int, mappings []mapping) int {
	currentValue := location
	for i := range mappings {
	currentMapping := mappings[len(mappings)-1-i]
	rangeLoop:
		for _, currentRange := range currentMapping.ranges {
			destinationFloor := currentRange.destinationRangeStart
			destinationCeiling := currentRange.destinationRangeStart + currentRange.rangeLength
			if currentValue >= destinationFloor && currentValue < destinationCeiling {
				diff := currentValue - destinationFloor
				currentValue = currentRange.sourceRangeStart + diff
				break rangeLoop
			}
		}
	}
	return currentValue
}

func min(numbers []int) int {
	minimum := math.MaxInt
	for _, number := range numbers {
		if number < minimum {
			minimum = number
		}
	}
	return minimum

}
