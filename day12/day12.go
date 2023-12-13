package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

type record struct {
	springs string
	manual []int
}

type cacheKey struct {
	i, blockIndex, currBlockSize int
}

func main() {
	file, err := os.Open("day12input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	records := make([]record, 0)
	for _, line := range lines {
		items := strings.Split(line, " ")
		manualParts := make([]int, 0)
		for _, value := range strings.Split(items[1], ",") {
			part, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("manual part is not a number")
			}
			manualParts = append(manualParts, part)
		}
		records = append(records, record{springs: items[0], manual:manualParts})
	}
	part2(records)
}

func part1(records []record) {
	total := 0
	for _, r := range records {
		total += numPermutations(0, r.springs, r.manual, 0, 0)
	}
	fmt.Println(total)
}

var cache = make(map[cacheKey]int)

func part2(records []record) {
	total := 0
	for _, r := range records {
		springs := strings.Repeat(r.springs+"?", 5)
		springs = springs[:len(springs)-1]
		manual := make([]int, 0)
		for i := 0; i < 5; i++ {
			manual = append(manual, r.manual...)
		}
		cache = make(map[cacheKey]int)
		value := numPermutations(0, springs, manual, 0, 0)
		fmt.Println(value)
		total += value
	}
	fmt.Println(total)
}


func numPermutations(i int, springs string, manual []int, blockIndex int, currBlockSize int) int {
	key := cacheKey{i:i, blockIndex:blockIndex, currBlockSize:currBlockSize}
	if i, ok := cache[key]; ok {
		return i
	}
	if currBlockSize > max(manual) {
		return 0
	}
	if i == len(springs) {
		if blockIndex == len(manual) && currBlockSize == 0 {
			return 1
		}
		if blockIndex == len(manual)-1 && manual[blockIndex] == currBlockSize {
			return 1
		}
		return 0
	}
	ans := 0
	for _, c := range [2]int{'.', '#'} {
		if springs[i] == byte(c) || springs[i] == '?' {
			if c == '.' && currBlockSize == 0 {
				ans += numPermutations(i+1, springs, manual, blockIndex, 0)
			} else if c == '.' && currBlockSize > 0 && blockIndex < len(manual) && manual[blockIndex] == currBlockSize {
				ans += numPermutations(i+1, springs, manual, blockIndex+1, 0)
			} else if c == '#' {
				ans += numPermutations(i+1, springs, manual, blockIndex, currBlockSize+1)
			}
		}
	}
	cache[key] = ans
	return ans
}

func blocksMatch(manual []int, blocks []int) bool {
	if len(blocks) > len(manual) {
		return false
	}
	mismatch := false
	for i := 0; i < len(blocks); i++ {
		if manual[i] != blocks[i] {
			mismatch = true
		}
	}
	return !mismatch
}

func isValid(springs string, manual []int) bool {
	block := 0
	blocks := make([]int, 0)
	for _, spring := range springs {
		if spring == '#' {
			block++
		} else if block != 0 && spring != '#' {
			blocks = append(blocks, block)
			block = 0
		}
	}
	if block != 0 {
		blocks = append(blocks, block)
	}
	if len(blocks) != len(manual) {
		return false
	}
	mismatch := false
	for i := range manual {
		if blocks[i] != manual[i] {
			mismatch = true
		}
	}
	return !mismatch
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