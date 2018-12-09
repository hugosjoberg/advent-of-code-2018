package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func time(c rune) int {
	return 60 + int(c) - int('A')
}

func main() {
	file, _ := os.Open("test.txt")
	scan := bufio.NewScanner(file)

	steps := make(map[rune][]rune)
	parents := make(map[rune]int)
	var first, second rune
	for scan.Scan() {
		line := scan.Text()
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &first, &second)
		steps[first] = append(steps[first], second)
		parents[second]++
	}
	for _, v := range steps {
		sort.Sort(sortRunes(v))
	}

	done := make([]rune, 0)
	for k := range steps {
		if parents[k] == 0 {
			done = append(done, k)
		}
	}

	answer := ""
	for len(done) > 0 {
		temp := make([]rune, len(done))
		copy(temp, done)
		sort.Sort(sortRunes(temp))
		x := temp[0]
		for i := 0; i < len(done); i++ {
			if done[i] == x {
				done = append(done[:i], done[i+1:]...)
			}
		}
		answer = answer + string(x)
		for _, v := range steps[x] {
			parents[v] = parents[v] - 1
			if parents[v] == 0 {
				done = append(done, v)
			}
		}
	}
	fmt.Println(answer)
	// Part 2

}
