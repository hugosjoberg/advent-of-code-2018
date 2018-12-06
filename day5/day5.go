package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func loopPolymer(polymer []rune) int {
	var newPolymer []rune
	for {
		newPolymer = reducePolymer(polymer)
		if string(polymer) == string(newPolymer) {
			return len(polymer)
		} else {
			polymer = newPolymer
		}
	}
	return 0
}

func removeUnit(polymer []rune, unit rune) []rune {
	for i := range polymer {

		if unicode.ToUpper(polymer[i]) == unit {
			polymer = append(polymer[:i], polymer[i+1:]...)
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

	polymer := []rune(s)
	fmt.Println(loopPolymer(polymer))

	s2 := s
	counter := len(polymer)
	for i := rune('A'); i <= rune('Z'); i++ {
		fmt.Println(string(i))
		s2 = s
		s2 = strings.Replace(s2, string(i), "", -1)
		s2 = strings.Replace(s2, string(unicode.ToLower(i)), "", -1)
		polymer = []rune(s2)
		lenPolymer := loopPolymer(polymer)
		fmt.Println(lenPolymer)
		if counter > lenPolymer {
			counter = lenPolymer
		}
	}
	fmt.Println(counter)

}
