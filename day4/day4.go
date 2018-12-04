package main

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type date struct {
	year, month, day, hour, minute int
	sleep                          bool
}

type guard struct {
	id int
	date
	sleepCounter map[int]int
}

func main() {
	file, _ := ioutil.ReadFile("input_day4.txt")
	// Solution to first part
	var schedule []string
	for _, line := range strings.Split(string(file), "\n") {
		schedule = append(schedule, line)
	}

	sort.Strings(schedule)
	var guards []guard
	id := 0
	for _, line := range schedule {
		var d date
		var g guard

		re, _ := regexp.Compile(`^\[([0-9]+)-`)
		match := re.FindStringSubmatch(line)
		d.year, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`-([0-9]+)-`)
		match = re.FindStringSubmatch(line)
		d.month, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`-([0-9]+) `)
		match = re.FindStringSubmatch(line)
		d.day, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(` ([0-9]+):`)
		match = re.FindStringSubmatch(line)
		d.hour, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`:([0-9]+)\]`)
		match = re.FindStringSubmatch(line)
		d.minute, _ = strconv.Atoi(match[1])

		if strings.Contains(line, "wakes up") && strings.Contains(line, "begins shift") {
			d.sleep = false
		} else {
			d.sleep = true
		}

		if strings.Contains(line, "begins shift") {
			re, _ = regexp.Compile(`#([0-9]+) `)
			match = re.FindStringSubmatch(line)
			g.id, _ = strconv.Atoi(match[1])
			id = g.id
		} else {
			g.id = id
		}
		guards.date = d
		guards = append(guards, g)
	}
}
