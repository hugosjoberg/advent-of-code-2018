package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type rect struct {
	x0, x1, y0, y1 int
	id             string
}

func main() {
	file, _ := ioutil.ReadFile("input_day3.txt")
	area := 0
	// Solution to first part
	var rectangles []rect
	areaMem := make(map[point]int)
	for _, line := range strings.Split(string(file), "\n") {
		var s point
		var r rect

		re, _ := regexp.Compile("^#([0-9]+)? @")
		match := re.FindStringSubmatch(line)
		r.id = match[1]

		re, _ = regexp.Compile("@ (.+)?,")
		match = re.FindStringSubmatch(line)
		padLeft, _ := strconv.Atoi(match[1])

		re, _ = regexp.Compile(",(.+)?:")
		match = re.FindStringSubmatch(line)
		padTop, _ := strconv.Atoi(match[1])

		re, _ = regexp.Compile(": (.+)?x")
		match = re.FindStringSubmatch(line)
		wide, _ := strconv.Atoi(match[1])

		re, _ = regexp.Compile("x([0-9]+)")
		match = re.FindStringSubmatch(line)
		tall, _ := strconv.Atoi(match[1])

		r.x0 = padLeft
		r.x1 = padLeft + wide
		r.y0 = padTop
		r.y1 = padTop + tall
		for y := r.y0; y < r.y1; y++ {
			for x := r.x0; x < r.x1; x++ {
				s.x = x
				s.y = y
				if _, exists := areaMem[s]; exists {
					areaMem[s]++
				} else {
					areaMem[s] = 1
				}
			}
		}
		rectangles = append(rectangles, r)

	}
	for _, v := range areaMem {
		if v > 1 {
			area++
		}
	}
	fmt.Println(area)

	for _, r := range rectangles {
		var s point
		flag := true
		for y := r.y0; y < r.y1; y++ {
			for x := r.x0; x < r.x1; x++ {
				s.x = x
				s.y = y
				if areaMem[s] > 1 {
					flag = false
				}
			}

		}
		if flag == true {
			fmt.Println(r.id)
		}

	}

}
