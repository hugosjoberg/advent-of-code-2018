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
	return abs(pc.x-c.x) + abs(pc.y-c.y)
}

func shortestDist(pc cord, cords map[int]cord) int {
	dist := 1000
	var tempDist, id int
	for i, c := range cords {
		if pc.x == c.x && pc.y == c.y {
			return i
		}
		tempDist = distance(pc, c)

		if tempDist == dist {
			id = 0
		} else if dist > tempDist {
			dist = tempDist
			id = i
		}
	}
	return id
}

func findInfinite(pc cord, minX, maxX, minY, maxY int) bool {
	if pc.x == minX || pc.x == maxX || pc.y == minY || pc.y == maxY {
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
	cords := make(map[int]cord)
	infinity := make(map[int]bool)
	var minX, minY, maxX, maxY = 1000, 1000, -1000, -1000
	for scan.Scan() {
		line := scan.Text()
		var co cord
		fmt.Sscanf(line, "%d, %d", &co.x, &co.y)
		cords[id] = co
		infinity[id] = false
		id++
		if co.x < minX {
			minX = co.x
		} else if co.x > maxX {
			maxX = co.x
		}
		if co.y < minY {
			minY = co.y
		} else if co.y > maxY {
			maxY = co.y
		}
	}

	area := make(map[int]int)
	for x := minX; x <= maxX; x++ {
		for y := -minY; y <= maxY; y++ {
			var co cord
			co.x = x
			co.y = y
			if x == 118 && y == 274 {
			}
			id = shortestDist(co, cords)
			area[id]++
			system[co] = id
			if infinity[id] == false {
				infinity[id] = findInfinite(co, minX, maxX, minY, maxY)
			}
		}
	}
	maxArea := 0
	for id, v := range area {
		if infinity[id] == false {
			if v > maxArea {
				maxArea = v
			}
		}
	}
	fmt.Println("Part1 :")
	fmt.Println(maxArea)

	m := make(map[cord]int)
	for x := minX; x <= maxX; x++ {
		for y := -minY; y <= maxY; y++ {
			sum := 0
			var c cord
			c.x = x
			c.y = y
			for _, co := range cords {
				sum += distance(c, co)
			}
			m[c] = sum
		}
	}
	maxArea = 0
	for _, d := range m {
		if d < 10000 {
			maxArea++
		}
	}
	fmt.Println(maxArea)
}
