package main

import (
	"bufio"
	"fmt"
	"os"
)

type cord struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(pc, c cord) int {
	dist := abs(pc.x-c.x) + abs(pc.y-c.y)
	return dist
}

func shortestDist(pc cord, cords map[cord]int) int {
	id := 0
	var dist, tempDist int
	for c, i := range cords {
		if pc.x == c.x && pc.y == c.y {
			return i
		}
		tempDist = distance(pc, c)
		if tempDist == 0 {
			return i
		} else if dist == 0 {
			dist = tempDist
			id = i

		} else if dist > tempDist {
			dist = tempDist
			id = i
		}
	}
	return id
}

func findInfinite(pc cord) bool {
	if pc.x == 0 || pc.x == 1000 {
		return true
	} else if pc.y == 0 || pc.y == 1000 {
		return true
	}
	return false
}

func calArea(system map[cord]int, id int) int {
	counter := 0
	for _, v := range system {
		if v == id {
			counter++
		}
	}
	return counter
}

func main() {
	file, _ := os.Open("input_day6.txt")
	scan := bufio.NewScanner(file)

	id := 1
	system := make(map[cord]int)
	cords := make(map[cord]int)
	infinity := make(map[int]bool)

	for scan.Scan() {
		line := scan.Text()
		var co cord
		fmt.Sscanf(line, "%d, %d", &co.x, &co.y)
		cords[co] = id
		infinity[id] = false
		id++
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			var co cord
			co.x = x
			co.y = y
			id = shortestDist(co, cords)
			system[co] = id
			if infinity[id] == false {
				infinity[id] = findInfinite(co)
			}
		}
	}
	area := 0
	for id, inf := range infinity {
		if inf == false {
			tempArea := calArea(system, id)
			if area > tempArea {
				area = tempArea
			} else if area == 0 {
				area = tempArea
			}
		}
	}
	fmt.Println(area)

}
