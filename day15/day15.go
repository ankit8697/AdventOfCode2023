package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
	"strings"
)

type lens struct {
	label string
	value int
}

type box struct {
	lenses []lens
}

func main() {
	file, err := os.Open("day15input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(contents), ",")
	part1(words)
	part2(words)
}

func part1(words []string) {
	ans := 0
	for _, word := range words {
		ans += hash(word)
	}
	fmt.Println(ans)
}

func part2(words []string) {
	var boxes [256]box
	for _, word := range words {
		if strings.Contains(word, "=") {
			parts := strings.Split(word, "=")
			label := parts[0]
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("lens value is not a number")
			}
			l := lens{label: label, value: value}
			boxNo := hash(label)
			boxes[boxNo].addLens(l)
		} else {
			label := word[:len(word)-1]
			boxNo := hash(label)
			boxes[boxNo].removeLens(label)
		}
	}
	ans := 0
	for i, box := range boxes {
		for j, l := range box.lenses {
			fp := (1+i)*(1+j)*l.value
			ans += fp
		}
	}
	fmt.Println(ans)
}

func hash(word string) int {
	sum := 0
	for i := range word {
		sum += int(word[i])
		sum *= 17
		sum = sum % 256
	}
	return sum
}

func (b *box) addLens(newLens lens) {
	if slices.ContainsFunc(b.lenses, func(l lens) bool {
		return newLens.label == l.label
	}) {
		i := slices.IndexFunc(b.lenses, func(l lens) bool {
			return newLens.label == l.label
		})
		b.lenses[i] = newLens
		return
	}
	b.lenses = add(b.lenses, newLens)
}

func (b *box) removeLens(label string) {
	if !slices.ContainsFunc(b.lenses, func(l lens) bool {
		return l.label == label
	}) {
		return
	}
	i := slices.IndexFunc(b.lenses, func(l lens) bool {
		return l.label == label
	})
	b.lenses = remove(b.lenses, i)
}

func remove(slice []lens, i int) []lens {
	return append(slice[:i], slice[i+1:]...)
}

func add(lenses []lens, l lens) []lens {
	return append(lenses, l)
}
