package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isUpper(r rune) bool {
	if r == unicode.ToUpper(r) {
		return true
	}
	return false
}

func reducePolymer(polymer []rune) []rune {
	var first, second bool
	var tempFirst, tempSecond rune
	var newPolymer []rune
	for i := range polymer {
		if i > len(polymer)-2 {
			return newPolymer
		}
		first, second = false, false
		if isUpper(polymer[i]) {
			first = true
			tempFirst = polymer[i]
		} else {
			tempFirst = unicode.ToUpper(polymer[i])
		}
		if isUpper(polymer[i+1]) {
			second = true
		} else {
			tempSecond = unicode.ToUpper(polymer[i+1])
		}
		if tempFirst == tempSecond && first != second {
			continue
		}
		newPolymer = append(newPolymer, tempFirst)
	}
	return newPolymer
}

func main() {
	file, _ := os.Open("input_day5.txt")
	scan := bufio.NewScanner(file)
	var s string
	for scan.Scan() {
		s = scan.Text()

	}
	var newPolymer, polymer []rune
	polymer = []rune(s)

	for {
		newPolymer = reducePolymer(polymer)
		if string(polymer) == string(newPolymer) {
			fmt.Println(polymer)
			break
		} else {
			polymer = newPolymer
		}
	}
}
