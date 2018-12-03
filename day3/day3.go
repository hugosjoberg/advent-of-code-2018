package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type information struct {
	id      string
	padLeft int
	padTop  int
	wide    int
	tall    int
}

func main() {
	file, _ := ioutil.ReadFile("input_day3.txt")
	// Solution to first part
	for _, line := range strings.Split(string(file), "\n") {
		re, _ := regexp.Compile("[^#](.+)? @")
		match := re.FindStringSubmatch(line)
		id := match[1]
		fmt.Println(id)

		re, _ = regexp.Compile("@ (.+)?,")
		match = re.FindStringSubmatch(line)
		padLeft := match[1]
		fmt.Println(padLeft)

		re, _ = regexp.Compile(",(.+)?:")
		match = re.FindStringSubmatch(line)
		padTop := match[1]
		fmt.Println(padTop)

		re, _ = regexp.Compile(": (.+)?x")
		match = re.FindStringSubmatch(line)
		wide := match[1]
		fmt.Println(wide)

		re, _ = regexp.Compile("x(.+)")
		match = re.FindStringSubmatch(line)
		tall := match[1]
		fmt.Println(tall)
	}

}
