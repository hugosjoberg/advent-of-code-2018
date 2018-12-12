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

func plantFinder(initial string, rules map[string]byte) string {
	var newString bytes.Buffer
	for i := 1; i < len(initial); i++ {
		temp := ""
		if i < 4 {
			temp = initial[0:i]
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

	initial := ""
	var rules = make(map[string]byte)
	for scan.Scan() {
		line := scan.Text()
		if len(line) > 10 {
			initial = line[15:]
		} else if len(line) > 0 {
			rules[line[0:5]] = line[len(line)-1]
		}
	}
	sum := 0
	fmt.Println(initial)
	for i := 0; i < 20; i++ {
		initial = plantFinder(initial, rules)
	}
	for j := 0; j < len(initial); j++ {
		if string(initial[j]) == "#" {
			sum = sum + j
		}
	}
	fmt.Println(sum)

}
