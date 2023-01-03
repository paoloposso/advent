package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// topThree := elvesTotals := sint{0,0,0}

	file, _ := ioutil.ReadFile("./input/test.txt")
	content := string(file)

	elves := strings.Split(content, "\n\n")

	biggest := 0

	for _, elf := range elves {
		total := 0

		itemsInBag := strings.Split(elf, "\n")

		for _, bag := range itemsInBag {
			c, _ := strconv.Atoi(bag)
			total += c
		}

		if total > biggest {
			biggest = total
		}
	}

	fmt.Println(biggest)
}
