package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func padLeft(plant string) string {
	var temp bytes.Buffer
	for i := 0; i < 5-len(plant); i++ {
		temp.WriteString(".")
	}
	temp.WriteString(plant)
	return temp.String()
}

func padRight(plant string) string {
	var temp bytes.Buffer
	temp.WriteString(plant)
	for i := 0; i < 5; i++ {
		if i >= len(plant) {
			temp.WriteString(".")
		}
	}
	return temp.String()
}

func matchRules(plant string, rules map[string]byte) string {
	for k, v := range rules {
		if k == plant {
			return string(v)
		}
	}
	return string('.')
}

func removePaddingLeft(plant string) int {
	left := 0
	for i := 0; i < len(plant); i++ {
		if string(plant[i]) == "." {
			left++
		} else if string(plant[i]) == "#" {
			break
		}
	}
	return left
}

func removePaddingRight(plant string) int {
	right := 0
	for i := len(plant) - 5; i < len(plant); i++ {
		if string(plant[i]) == "." {
			right++
		} else if string(plant[i]) == "#" {
			return 0
		}
	}
	return right
}

func plantFinder(initial string, rules map[string]byte) string {
	var newString bytes.Buffer
	for i := -5; i < len(initial); i++ {
		temp := ""
		if i < 0 {
			temp = initial[0 : 5+i]
			temp = padLeft(temp)
		} else if i > len(initial)-5 {
			temp = initial[i:]
			temp = padRight(temp)
		} else {
			temp = initial[i : i+5]
		}
		newString.WriteString(matchRules(temp, rules))

	}
	return newString.String()
}

func main() {
	file, _ := os.Open("input.txt")
	scan := bufio.NewScanner(file)

	initialOriginal := ""
	var rules = make(map[string]byte)
	for scan.Scan() {
		line := scan.Text()
		if len(line) > 10 {
			initialOriginal = line[15:]
		} else if len(line) > 0 {
			rules[line[0:5]] = line[len(line)-1]
		}
	}
	// Part 1
	sum := 0
	leftStart := 0
	initial := initialOriginal
	for i := 0; i < 20; i++ {
		initial = plantFinder(initial, rules)
		paddingRight := removePaddingRight(initial)
		paddingLeft := removePaddingLeft(initial)
		initial = initial[paddingLeft : len(initial)-paddingRight]
		leftStart = paddingLeft - 5
	}

	for j := 0; j < len(initial); j++ {
		if string(initial[j]) == "#" {
			sum = sum + j + leftStart
		}
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	currentSum := 0
	average := 0
	leftStart = 0
	initial = initialOriginal
	for i := 0; i < 50000000000; i++ {
		initial = plantFinder(initial, rules)
		paddingRight := removePaddingRight(initial)
		paddingLeft := removePaddingLeft(initial)
		initial = initial[paddingLeft : len(initial)-paddingRight]
		leftStart = paddingLeft - 5
		if i >= 2000 && i < 2100 {
			for j := 0; j < len(initial); j++ {
				if string(initial[j]) == "#" {
					sum = sum + j + leftStart
				}
			}
		} else if i == 2100 {
			average = sum / 100
			break
		}
		for j := 0; j < len(initial); j++ {
			if string(initial[j]) == "#" {
				currentSum = currentSum + j + leftStart
			}
		}

	}
	fmt.Println(((50000000000 - 2100) * average) + currentSum)
}
