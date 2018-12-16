package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	recipes := "825401"
	scores := []byte{'3', '7'}
	elf1, elf2 := 0, 1
	for len(scores) < 50000000 {
		tempElf1, _ := strconv.Atoi(string(scores[elf1]))
		tempElf2, _ := strconv.Atoi(string(scores[elf2]))
		scores = append(scores, strconv.Itoa(tempElf1+tempElf2)...)
		elf1 = (elf1 + tempElf1 + 1) % len(scores)
		elf2 = (elf2 + tempElf2 + 1) % len(scores)
	}
	tempRecipes, _ := strconv.Atoi(recipes)
	fmt.Println("Part 1:", string(scores[tempRecipes:tempRecipes+10]))
	fmt.Println("Part 2:", strings.Index(string(scores), recipes))
}
