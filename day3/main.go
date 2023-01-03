package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var arr = [...]string{".", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func getPriority(ch string) int {
	isLowercase := strings.ToLower(ch) == ch

	for ix, item := range arr {
		if strings.ToUpper(ch) == item {
			if !isLowercase {
				return ix + 26
			}
			return ix
		}
	}
	panic(fmt.Sprint("invalid char"))
}

func main() {
	fileBytes, _ := ioutil.ReadFile("./input/input2.txt")
	sacks := strings.Split(string(fileBytes), "\n")

	result := 0

	for i := 0; i < len(sacks)-2; i = i + 3 {
		groupOfSacks := sacks[i : i+3]

		for _, itemOnSack1 := range groupOfSacks[0] {
			repeatedItem := sackContains(groupOfSacks[1], string(itemOnSack1))
			repeatedItem = sackContains(groupOfSacks[2], repeatedItem)
			if repeatedItem != "" {
				repeatedItem = string(itemOnSack1)

				r := getPriority(repeatedItem)
				fmt.Println(r)
				result += r

				break
			}
		}
	}

	fmt.Println(result)
}

func main1() {
	fileBytes, _ := ioutil.ReadFile("./input/input2.txt")
	sacks := strings.Split(string(fileBytes), "\n")

	result := 0

	for _, sack := range sacks {
		mid := (len(sack) / 2)
		sack1 := sack[:mid]
		sack2 := sack[mid:]

		for _, itemOnSack1 := range sack1 {
			repeatedItem := sackContains(sack2, string(itemOnSack1))
			if repeatedItem != "" {
				repeatedItem = string(itemOnSack1)

				// if sackContains(sack1, strings.ToUpper(repeatedItem)) != "" || sackContains(sack2, strings.ToUpper(repeatedItem)) != "" {
				// 	repeatedItem = strings.ToUpper(repeatedItem)
				// }

				r := getPriority(repeatedItem)
				fmt.Println(r)
				result += r

				break
			}
		}

	}
	fmt.Println(result)
}

func sackContains(sackItem string, s string) string {
	for _, item := range sackItem {
		if s == string(item) {
			return string(item)
		}
	}
	return ""
}
