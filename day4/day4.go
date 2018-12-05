package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type date struct {
	time  []time.Time
	sleep bool
}
type guard struct {
	date         time.Time
	id           int
	sleep        string
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

		var g guard
		var year, month, day, hour, minute int

		re, _ := regexp.Compile(`^\[([0-9]+)-`)
		match := re.FindStringSubmatch(line)
		year, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`-([0-9]+)-`)
		match = re.FindStringSubmatch(line)
		month, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`-([0-9]+) `)
		match = re.FindStringSubmatch(line)
		day, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(` ([0-9]+):`)
		match = re.FindStringSubmatch(line)
		hour, _ = strconv.Atoi(match[1])

		re, _ = regexp.Compile(`:([0-9]+)\]`)
		match = re.FindStringSubmatch(line)
		minute, _ = strconv.Atoi(match[1])
		g.date = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

		if strings.Contains(line, "falls asleep") {
			g.sleep = "sleeps"
		} else if strings.Contains(line, "wakes up") {
			g.sleep = "wakes"
		}

		if strings.Contains(line, "begins shift") {
			re, _ = regexp.Compile(`#([0-9]+) `)
			match = re.FindStringSubmatch(line)
			id, _ = strconv.Atoi(match[1])
			g.id = id
			g.sleep = "begins"
			guards = append(guards, g)
		} else {
			g.id = id
			guards = append(guards, g)
		}
	}

	asleep := map[int]int{}
	var sleepiestGuard, from, guardID int
	for _, g := range guards {
		switch g.sleep {
		case "begins":
			guardID = g.id
		case "sleeps":
			from = g.date.Minute()
		case "wakes":
			t := g.date.Minute() - from
			asleep[guardID] += t
			if asleep[guardID] > asleep[sleepiestGuard] {
				sleepiestGuard = guardID
			}
		}
	}
	fmt.Println(sleepiestGuard)

	minutes := [60]int{}
	guardID = -1
	var sleepyminute int
	for _, g := range guards {
		if g.sleep == "begins" {
			guardID = g.id
			continue
		}
		if guardID != sleepiestGuard {
			continue
		}
		switch g.sleep {
		case "sleeps":
			from = g.date.Minute()
		case "wakes":
			to := g.date.Minute()
			for i := from; i < to; i++ {
				minutes[i]++
				if minutes[i] > minutes[sleepyminute] {
					sleepyminute = i
				}
			}
		}
	}

	fmt.Println(sleepiestGuard * sleepyminute)

	minute := map[int]*[60]int{}
	guardID = -1
	sleepyminute, sleepiestGuard = 0, 0
	for _, g := range guards {
		switch g.sleep {
		case "begins":
			guardID = g.id
			if minute[guardID] == nil {
				minute[guardID] = &[60]int{}
			}
			if minute[sleepiestGuard] == nil {
				sleepiestGuard = guardID
			}
		case "sleeps":
			from = g.date.Minute()
		case "wakes":
			to := g.date.Minute()
			for i := from; i < to; i++ {
				minute[guardID][i]++
				if minute[guardID][i] > minute[sleepiestGuard][sleepyminute] {
					sleepiestGuard = guardID
					sleepyminute = i
				}
			}
		}
	}
	fmt.Println(sleepiestGuard * sleepyminute)
}
