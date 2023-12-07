package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bid   int
}

type HandType int32

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("day7input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	hands := make([]hand, 0)
	for _, line := range lines {
		values := strings.Fields(line)
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("bid is not a number")
		}
		hands = append(hands, hand{cards: values[0], bid: bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})
	total := 0
	for i, hand := range hands {
		total = total + (i+1)*hand.bid
	}
	fmt.Println(total)
}

func part2() {
	file, err := os.Open("day7input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	hands := make([]hand, 0)
	for _, line := range lines {
		values := strings.Fields(line)
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("bid is not a number")
		}
		hands = append(hands, hand{cards: values[0], bid: bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHandsWithJokers(hands[i], hands[j])
	})
	total := 0
	for i, hand := range hands {
		total = total + (i+1)*hand.bid
	}
	fmt.Println(total)
}

func categorizeHand(hand string) HandType {
	handComposition := make(map[string]int)
	for _, card := range hand {
		handComposition[string(card)]++
	}
	if len(handComposition) == 1 {
		return FiveOfAKind
	} else if len(handComposition) == 2 {
		for _, value := range handComposition {
			if value == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
		// Could be three of a kind or two pair
	} else if len(handComposition) == 3 {
		for _, value := range handComposition {
			if value == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	} else if len(handComposition) == 4 {
		return OnePair
	}
	return HighCard
}

func categorizeHandWithJokers(hand string) HandType {
	handComposition := make(map[string]int)
	for _, card := range hand {
		handComposition[string(card)]++
	}
	if len(handComposition) == 1 {
		return FiveOfAKind
	} else if len(handComposition) == 2 {
		if _, hasJack := handComposition["J"]; hasJack {
			return FiveOfAKind
		}
		for _, value := range handComposition {
			if value == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
		// Could be three of a kind or two pair
	} else if len(handComposition) == 3 {
		numJack, hasJack := handComposition["J"]
		for _, value := range handComposition {
			if value == 3 && hasJack {
				return FourOfAKind
			} else if value == 3 && !hasJack {
				return ThreeOfAKind
			}
		}
		if hasJack {
			if numJack == 1 {
				return FullHouse
			}
			return FourOfAKind
		}
		return TwoPair
	} else if len(handComposition) == 4 {
		if _, hasJack := handComposition["J"]; hasJack {
			return ThreeOfAKind
		}
		return OnePair
	}
	if _, hasJack := handComposition["J"]; hasJack {
		return OnePair
	}
	return HighCard
}

func compareHands(hand1 hand, hand2 hand) bool {
	hand1Type := categorizeHand(hand1.cards)
	hand2Type := categorizeHand(hand2.cards)
	if hand1Type != hand2Type {
		return hand1Type < hand2Type
	}
	cardName := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	for i := range hand1.cards {
		if hand1.cards[i] == hand2.cards[i] {
			continue
		}
		return cardName[string(hand1.cards[i])] < cardName[string(hand2.cards[i])]
	}
	return false
}

func compareHandsWithJokers(hand1 hand, hand2 hand) bool {
	hand1Type := categorizeHandWithJokers(hand1.cards)
	hand2Type := categorizeHandWithJokers(hand2.cards)
	if hand1Type != hand2Type {
		return hand1Type < hand2Type
	}
	cardName := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}
	for i := range hand1.cards {
		if hand1.cards[i] == hand2.cards[i] {
			continue
		}
		return cardName[string(hand1.cards[i])] < cardName[string(hand2.cards[i])]
	}
	return false
}
