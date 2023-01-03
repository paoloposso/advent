package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("./input/input.txt")
	allSectionPairs := strings.Split(string(fileBytes), "\n")

	totalResult := 0

	for _, pair := range allSectionPairs {
		sections := strings.Split(pair, ",")

		fmt.Println(sections)

		if fullyContains(sections) {
			totalResult++
		}
	}

	fmt.Println(totalResult)
}

func fullyContains(sections []string) bool {
	leftSection := sections[0]
	rightSection := sections[1]

	leftSectionStart, err := strconv.Atoi(strings.Split(leftSection, "-")[0])
	leftSectionEnd, err := strconv.Atoi(strings.Split(leftSection, "-")[1])

	rightSectionStart, err := strconv.Atoi(strings.Split(rightSection, "-")[0])
	rightSectionEnd, err := strconv.Atoi(strings.Split(rightSection, "-")[1])

	if err != nil {
		panic("invalid range")
	}

	if leftSectionStart <= rightSectionStart && leftSectionEnd >= rightSectionEnd {
		return true
	} else if leftSectionStart >= rightSectionStart && rightSectionEnd >= leftSectionEnd {
		return true
	} else {
		return false
	}
}
