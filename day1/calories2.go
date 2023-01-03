package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func mainx() {

	elvesTotals := []int{0, 0, 0}

	file, _ := ioutil.ReadFile("./input/test.txt")
	content := string(file)

	elves := strings.Split(content, "\n\n")

	for _, elf := range elves {
		total := 0

		itemsInBag := strings.Split(elf, "\n")

		for _, bag := range itemsInBag {
			c, _ := strconv.Atoi(bag)
			total += c
		}
		elvesTotals = append(elvesTotals, total)
	}

	sort.Ints(elvesTotals)

	fmt.Println(elvesTotals)
}
