package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day1input.txt")
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

	total := 0
	for _, v := range strings.Split(string(contents), "\n") {
		number, err := getCode(v)
		if err != nil {
			fmt.Printf("Could not get code for string: %s\n", v)
		}
		total += number
	}
	fmt.Println(total)

	totalPart2 := 0
	for _, v := range strings.Split(string(contents), "\n") {
		number, err := getCodeForPartTwo(v)
		if err != nil {
			fmt.Printf("Could not get code for string: %s\n", v)
		}
		totalPart2 += number
	}
	fmt.Println(totalPart2)
	// getCodeForPartTwo("eightwo")
}

func getCode(word string) (int, error) {
	numbers := ""
	for _, v := range word {
		if unicode.IsDigit(v) {
			numbers = numbers + string(v)
		}
	}
	if len(numbers) == 1 {
		numbers = numbers + numbers
	}
	finalNumber := string(numbers[0]) + string(numbers[len(numbers)-1])
	return strconv.Atoi(string(finalNumber))
}

func getCodeForPartTwo(word string) (int, error) {
	var numberMap = make(map[string]string)
	numberMap["one"] = "1"
	numberMap["two"] = "2"
	numberMap["three"] = "3"
	numberMap["four"] = "4"
	numberMap["five"] = "5"
	numberMap["six"] = "6"
	numberMap["seven"] = "7"
	numberMap["eight"] = "8"
	numberMap["nine"] = "9"

	potentialNumbers := make([]string, 10)
	for i := range word {
		potentialNumbers = append(potentialNumbers, getNumForWord(word[i:])...)
	}
	var cleanedNumbers []string
	for _, number := range potentialNumbers {
		if number != "" {
			cleanedNumbers = append(cleanedNumbers, number)
		}
	}
	num1 := ""
	num2 := ""
	if value, ok := numberMap[cleanedNumbers[0]]; ok {
		num1 = value
	} else {
		num1 = cleanedNumbers[0]
	}
	if value, ok := numberMap[cleanedNumbers[len(cleanedNumbers)-1]]; ok {
		num2 = value
	} else {
		num2 = cleanedNumbers[len(cleanedNumbers)-1]
	}
	return strconv.Atoi(num1 + num2)
}

func getNumForWord(word string) []string {
	regexString := "one|two|three|four|five|six|seven|eight|nine|\\d"

	re := regexp.MustCompile(regexString)
	numbers := re.FindAllString(word, -1)
	return numbers
}
