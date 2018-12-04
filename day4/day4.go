package main

import (
	"fmt"
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
	dates        []date
	sleepCounter map[int]int
	totalSleep   int
}

func main() {
	file, _ := ioutil.ReadFile("input_day4.txt")
	// Solution to first part
	var schedule []string
	for _, line := range strings.Split(string(file), "\n") {
		schedule = append(schedule, line)
	}

	sort.Strings(schedule)

	guards := make(map[int]guard)

	id := 0
	for _, line := range schedule {
		var d date
		var g guard

		if strings.Contains(line, "falls asleep") {
			d.sleep = true
		} else {
			d.sleep = false
		}

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

		if strings.Contains(line, "begins shift") {
			re, _ = regexp.Compile(`#([0-9]+) `)
			match = re.FindStringSubmatch(line)
			id, _ = strconv.Atoi(match[1])
			if g.sleepCounter == nil {
				g.sleepCounter = make(map[int]int)
			}
			g.sleepCounter[d.minute] = 0
			g.totalSleep = 0
			g.dates = append(g.dates, d)
			guards[id] = g
		} else {
			if _, exists := guards[id].sleepCounter[d.minute]; exists {
				guards[id].sleepCounter[d.minute]++
			} else {
				guards[id].sleepCounter[d.minute] = 1
			}
			g = guards[id]
			g.totalSleep++

			g.dates = append(guards[id].dates, d)
			guards[id] = g
		}

	}

	for _, guard := range guards {
		for _, dat := range guard.dates {
			var d date
			d.year = dat.year
			d.month = dat.month
			d.day = dat.day
			d.hour = dat.hour
			d.minute = dat.minute + 1
			if dat.sleep == false {
				d.sleep = false
			} else {
				d.sleep = true
				if _, exists := guard.sleepCounter[d.minute]; exists {
					guard.sleepCounter[d.minute]++
					guard.totalSleep++
				}
			}
		}
	}
	sleepiestGuard := 0
	for key, guard := range guards {
		if sleepiestGuard == 0 {
			sleepiestGuard = key
		}
		if guard.totalSleep > guards[sleepiestGuard].totalSleep {
			sleepiestGuard = key
		}
	}
	fmt.Println(sleepiestGuard)
}
