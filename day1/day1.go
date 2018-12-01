package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input_day1.txt")
	answer := 0
	var freqList []int
	// Solution to first part
	for _, line := range strings.Split(string(file), "\n") {
		num, _ := strconv.Atoi(line)
		answer += num
		freqList = append(freqList, num)
	}
	fmt.Println(answer)
	// Solution to second part
	frequency := 0
	freqMemory := make(map[int]int)
	freqMemory[0] = 0
	for {
		for _, num := range freqList {
			frequency += num
			if _, exists := freqMemory[frequency]; exists {
				fmt.Println(frequency)
				os.Exit(0)
			}
			freqMemory[frequency] = 0
		}
	}
}
