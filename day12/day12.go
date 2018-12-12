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
	for i := 0; i < 5; i++ {
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
	for i := 0; i < len(initial)-5; i++ {
		temp := initial[i : i+5]
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
	paddingLeft := 0
	for i := 0; i < 20; i++ {
		initial = "...." + initial + "...."
		initial = plantFinder(initial, rules)
		paddingRight := removePaddingRight(initial)
		paddingLeft = removePaddingLeft(initial)
		initial = initial[paddingLeft : len(initial)-paddingRight]
		leftStart = paddingLeft + leftStart - 2
	}

	for j := 0; j < len(initial); j++ {
		if string(initial[j]) == "#" {
			sum = sum + j + leftStart
		}
	}
	fmt.Println(sum)

	// Part 2
	// Part 1
	sum = 0
	sum2 := 0
	diff := 0
	diff2 := 0
	leftStart = 0
	initial = initialOriginal
	paddingLeft = 0
	for i := 0; i < 100; i++ {
		initial = "...." + initial + "...."
		initial = plantFinder(initial, rules)
		paddingRight := removePaddingRight(initial)
		paddingLeft = removePaddingLeft(initial)
		initial = initial[paddingLeft : len(initial)-paddingRight]
		leftStart = paddingLeft + leftStart - 2
		for j := 0; j < len(initial); j++ {
			if string(initial[j]) == "#" {
				sum = sum + j + leftStart
			}
		}
		diff = sum - sum2
		//fmt.Println(diff - diff2)
		diff2 = sum - sum2
		sum2 = sum
	}
	fmt.Println(diff)
	fmt.Println(diff2)
	fmt.Println(sum)
	fmt.Println(((50000000000 - 100) * 20) + sum)
}

//fmt.Println(((50000000000 - 100) * 20) + currentSum)
