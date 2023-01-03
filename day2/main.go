package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	file, _ := ioutil.ReadFile("./input/test.txt")

	rounds := strings.Split(string(file), "\n")

	totalScore := 0

	for _, round := range rounds {

		roundScore := 0

		// A - X: Rock
		// B - Y: Paper
		// C - Z: Scissors

		r := strings.Split(round, " ")
		theirs := r[0]
		mine := r[1]

		if mine == "X" { // mine is rock
			roundScore = 1
			if theirs == "A" {
				roundScore += 3
			} else if theirs == "B" {
				roundScore += 0
			} else {
				roundScore += 6
			}
		} else if mine == "Y" {
			roundScore = 2
			if theirs == "A" {
				roundScore += 6
			} else if theirs == "B" {
				roundScore += 3
			} else {
				roundScore += 0
			}
		} else {
			roundScore = 3
			if theirs == "A" {
				roundScore += 0
			} else if theirs == "B" {
				roundScore += 6
			} else {
				roundScore += 3
			}
		}

		totalScore += roundScore
		fmt.Println(roundScore)
	}

	fmt.Println("total score: ", totalScore)
}

func main() {
	file, _ := ioutil.ReadFile("./input/test1.txt")

	rounds := strings.Split(string(file), "\n")

	totalScore := 0

	for _, round := range rounds {

		roundScore := 0

		// A - X: Rock
		// B - Y: Paper
		// C - Z: Scissors

		r := strings.Split(round, " ")
		theirs := r[0]
		condition := r[1]

		mine := ""

		if condition == "X" { //lose
			if theirs == "A" {
				mine = "Z"
			} else if theirs == "B" {
				mine = "X"
			} else if theirs == "C" {
				mine = "Y"
			}
		} else if condition == "Y" { //draw
			if theirs == "A" {
				mine = "X"
			} else if theirs == "B" {
				mine = "Y"
			} else if theirs == "C" {
				mine = "Z"
			}
		} else if condition == "Z" { //win
			if theirs == "A" {
				mine = "Y"
			} else if theirs == "B" {
				mine = "Z"
			} else if theirs == "C" {
				mine = "X"
			}
		} else {
			panic(fmt.Sprint("invalid input: ", condition))
		}

		if mine == "X" { // mine is rock
			roundScore = 1
			if theirs == "A" {
				roundScore += 3
			} else if theirs == "B" {
				roundScore += 0
			} else {
				roundScore += 6
			}
		} else if mine == "Y" {
			roundScore = 2
			if theirs == "A" {
				roundScore += 6
			} else if theirs == "B" {
				roundScore += 3
			} else {
				roundScore += 0
			}
		} else {
			roundScore = 3
			if theirs == "A" {
				roundScore += 0
			} else if theirs == "B" {
				roundScore += 6
			} else {
				roundScore += 3
			}
		}

		totalScore += roundScore
		fmt.Println(roundScore)
	}

	fmt.Println("total score: ", totalScore)
}
