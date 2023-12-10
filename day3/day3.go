package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	value  string
	row    int
	column int
}

type gearLocation struct {
	row    int
	column int
}

func main() {
	part1()
}

func part1() {
	file, err := os.Open("day3input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(contents), "\n")
	var numbers []number
	for i := range rows {
		numberBuffer := ""
		for j := range rows[i] {
			if unicode.IsDigit(rune(rows[i][j])) {
				numberBuffer += string(rows[i][j])
			} else if numberBuffer != "" {
				numbers = append(numbers, number{value: numberBuffer, row: i, column: j - 1})
				numberBuffer = ""
			}
		}
		if numberBuffer != "" {
			numbers = append(numbers, number{value: numberBuffer, row: i, column: len(rows[i]) - 1})
			numberBuffer = ""
		}
	}

	gearLocations := make(map[gearLocation][]string)

	total := 0
	for _, number := range numbers {
		row := number.row
		startCol := number.column - len(number.value) + 1

		// Check top left corner for symbol
		if row != 0 && startCol != 0 && isSymbol(rows[row-1][startCol-1]) {
			if string(rows[row-1][startCol-1]) == "*" {
				newGear := gearLocation{row: row - 1, column: startCol - 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check top right for symbol
		if row != 0 && number.column != len(rows[0])-1 && isSymbol(rows[row-1][number.column+1]) {
			if string(rows[row-1][number.column+1]) == "*" {
				newGear := gearLocation{row: row - 1, column: number.column + 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check bottom left corner for symbol
		if row != len(rows)-1 && startCol != 0 && isSymbol(rows[row+1][startCol-1]) {
			if string(rows[row+1][startCol-1]) == "*" {
				newGear := gearLocation{row: row + 1, column: startCol - 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check bottom right corner for symbol
		if row != len(rows)-1 && number.column != len(rows[0])-1 && isSymbol(rows[row+1][number.column+1]) {
			if string(rows[row+1][number.column+1]) == "*" {
				newGear := gearLocation{row: row + 1, column: number.column + 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check left for symbol
		if startCol != 0 && isSymbol(rows[row][startCol-1]) {
			if string(rows[row][startCol-1]) == "*" {
				newGear := gearLocation{row: row, column: startCol - 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check right for symbol
		if number.column != len(rows[0])-1 && isSymbol(rows[row][number.column+1]) {
			if string(rows[row][number.column+1]) == "*" {
				newGear := gearLocation{row: row, column: number.column + 1}
				gearLocations[newGear] = append(gearLocations[newGear], number.value)
			}
			total += getNumberValue(number)
			continue
		}
		// Check above number for symbol
		if row != 0 {
			foundSymbol := false
			for i := range number.value {
				if isSymbol(rows[row-1][number.column-i]) {
					if string(rows[row-1][number.column-i]) == "*" {
						newGear := gearLocation{row: row - 1, column: number.column - i}
						gearLocations[newGear] = append(gearLocations[newGear], number.value)
					}
					foundSymbol = true
				}
			}
			if foundSymbol {
				total += getNumberValue(number)
			}
		}
		// Check below number for symbol
		if row != len(rows)-1 {
			foundSymbol := false
			for i := range number.value {
				if isSymbol(rows[row+1][number.column-i]) {
					if string(rows[row+1][number.column-i]) == "*" {
						newGear := gearLocation{row: row + 1, column: number.column - i}
						gearLocations[newGear] = append(gearLocations[newGear], number.value)
					}
					foundSymbol = true
				}
			}
			if foundSymbol {
				total += getNumberValue(number)
			}
		}
	}
	fmt.Println(total)
	gearRatio := 0

	for _, value := range gearLocations {
		if len(value) != 2 {
			continue
		}
		firstNumber, err := strconv.Atoi(value[0])
		if err != nil {
			fmt.Println("first number is not a number")
		}
		secondNumber, err := strconv.Atoi(value[1])
		if err != nil {
			fmt.Println("second number is not a number")
		}
		gearRatio += firstNumber * secondNumber
	}
	fmt.Printf("gear ratio = %d\n", gearRatio)
}

func isSymbol(char byte) bool {
	return string(char) != "." && !unicode.IsDigit(rune(char))
}

func getNumberValue(number number) int {
	numberValue, err := strconv.Atoi(number.value)
	if err != nil {
		fmt.Println("number could not be made into int")
	}
	return numberValue
}
