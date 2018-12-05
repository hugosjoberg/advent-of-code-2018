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
	for i := range polymer {
		if i > len(polymer)-2 {
			return polymer
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
			tempSecond = polymer[i+1]
		} else {
			tempSecond = unicode.ToUpper(polymer[i+1])
		}

		if tempFirst == tempSecond && first != second {
			polymer = append(polymer[:i], polymer[i+2:]...)
			break
		}

	}
	return polymer
}

func removeUnit(polymer []rune, unit rune) []rune {
	var first, second bool
	var tempFirst, tempSecond rune
	for i := range polymer {
		if i > len(polymer)-2 {
			return polymer
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
			tempSecond = polymer[i+1]
		} else {
			tempSecond = unicode.ToUpper(polymer[i+1])
		}

		if tempFirst != unit && tempSecond != unit {
			continue
		}

		if tempFirst == tempSecond && first != second {
			polymer = append(polymer[:i], polymer[i+2:]...)
			break
		}

	}
	return polymer
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
	counter := len(polymer)
	for {
		newPolymer = reducePolymer(polymer)
		if string(polymer) == string(newPolymer) {
			fmt.Println(len(polymer))
			break
		} else {
			polymer = newPolymer
		}
	}
	for i := rune('A'); i <= rune('Z'); i++ {
		polymer = []rune(s)
		fmt.Println(string(i))
		for {
			newPolymer = removeUnit(polymer, rune(i))
			if string(polymer) == string(newPolymer) {
				if len(polymer) < counter {
					counter = len(polymer)
					fmt.Println(counter)
				}
				break
			} else {
				polymer = newPolymer
			}
		}
	}
	fmt.Println(counter)

}
