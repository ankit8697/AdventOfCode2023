package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day6input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	tempTime := ""
	for _, currentTime := range strings.Fields(lines[0])[1:] {
		tempTime += currentTime
	}
	tempDistance := ""
	for _, currentDistance := range strings.Fields(lines[1])[1:] {
		tempDistance += currentDistance
	}
	time, err := strconv.Atoi(tempTime)
	if err != nil {
		fmt.Println("time is not a number")
	}
	distance, err := strconv.Atoi(tempDistance)
	if err != nil {
		fmt.Println("distance is not a number")
	}
	wins := 0
	for chargeTime := 0; chargeTime <= time; chargeTime++ {
		distanceTravelled := (time - chargeTime) * chargeTime
		if distanceTravelled > distance {
			wins++
		}
	}
	fmt.Println(wins)
}

func part1() {
	file, err := os.Open("day6input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	times := make([]int, 0)
	for _, time := range strings.Fields(lines[0])[1:] {
		newTime, err := strconv.Atoi(time)
		if err != nil {
			fmt.Println("time is not a number")
		}
		times = append(times, newTime)
	}
	distances := make([]int, 0)
	for _, distance := range strings.Fields(lines[1])[1:] {
		newDistance, err := strconv.Atoi(distance)
		if err != nil {
			fmt.Println("distance is not a number")
		}
		distances = append(distances, newDistance)
	}
	product := 1
	for i := range times {
		time := times[i]
		distance := distances[i]
		wins := 0
		for chargeTime := 0; chargeTime <= time; chargeTime++ {
			distanceTravelled := (time - chargeTime) * chargeTime
			if distanceTravelled > distance {
				wins++
			}
		}
		if wins != 0 {
			product *= wins
		}
	}
	fmt.Println(product)
}
