package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calculateSum(slice []int) int {
	sum := 0
	for _, i := range slice {
		sum += i
	}
	return sum
}

func newScore(data, scores []int) int {
	sum := 0
	for k := range data {
		if k > 0 && k <= len(scores) {
			sum += scores[k-1]
		}
	}
	return sum
}

func findMeta(data []int) (int, int, []int) {
	children := data[0]
	metas := data[1]
	data = data[2:]
	var scores []int
	totals := 0

	var total, score int
	for i := 0; i < children; i++ {
		total, score, data = findMeta(data)
		totals += total
		scores = append(scores, score)
	}

	totals += calculateSum(data[:metas])

	if children == 0 {
		return totals, calculateSum(data[:metas]), data[metas:]
	}
	return totals, newScore(data[:metas], scores), data[metas:]

}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	var ints []int

	for _, d := range strings.Split(strings.TrimSpace(string(f)), " ") {
		i, _ := strconv.Atoi(d)
		ints = append(ints, i)
	}
	total, score, _ := findMeta(ints)
	fmt.Println(total)
	fmt.Println(score)
}
