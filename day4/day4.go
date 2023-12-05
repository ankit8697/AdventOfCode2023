package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	part2()
}

func part1() {
	cards := getCards()
	totalPoints := 0
	for _, card := range cards {
		totalPoints += getPointsForCard(card)
	}
	fmt.Println(totalPoints)
}

func part2() {
	cards := getCards()
	cardWins := make(map[string]int)
	cardCopies := make(map[string]int)
	for _, card := range cards {
		cardNumber := strings.Fields(strings.Split(card, ":")[0])
		cardWins[cardNumber[1]] = getNumWinsForCard(card)
		cardCopies[cardNumber[1]] = 1
	}
	for _, card := range cards {
		cardNumber := strings.Fields(strings.Split(card, ":")[0])[1]
		wins := cardWins[cardNumber]
		for i := 0; i < cardCopies[cardNumber]; i++ {
			cardNumberInt, err := strconv.Atoi(cardNumber)
			if err != nil {
				fmt.Println("Card number is not a number")
			}
			for i := cardNumberInt + 1; i <= cardNumberInt + wins; i++ {
				cardCopies[strconv.Itoa(i)]++
			}
		}
		
	}
	totalCopies := 0
	for _, value := range cardCopies {
		totalCopies += value
	}
	fmt.Println(totalCopies)
}

func getCards() []string {
	file, err := os.Open("day4input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	cards := strings.Split(string(contents), "\n")
	return cards
}

func getPointsForCard(card string) int {
	splitCard := strings.Split(card, ":")
		splitNumbers := strings.Split(splitCard[1], "|")
		winningNumbers := strings.Split(splitNumbers[0], " ")
		yourNumbers := strings.Split(splitNumbers[1], " ")
		winningSet := make(map[string]bool)
		for _, number := range winningNumbers {
			winningSet[number] = true
		}
		delete(winningSet, "")
		isFirstMatch := true
		points := 0
		for _, number := range yourNumbers {
			if _, ok := winningSet[number]; ok {
				if isFirstMatch {
					points = 1
					isFirstMatch = false
				} else {
					points *= 2
				}
			}
		}
		return points
}

func getNumWinsForCard(card string) int {
	splitCard := strings.Split(card, ":")
		splitNumbers := strings.Split(splitCard[1], "|")
		winningNumbers := strings.Split(splitNumbers[0], " ")
		yourNumbers := strings.Split(splitNumbers[1], " ")
		winningSet := make(map[string]bool)
		for _, number := range winningNumbers {
			winningSet[number] = true
		}
		delete(winningSet, "")
		wins := 0
		for _, number := range yourNumbers {
			if _, ok := winningSet[number]; ok {
				wins++
			}
		}
		return wins
}