package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x, y, velX, velY int
}

func tic(cords []coord) []coord {
	for i, c := range cords {
		c.x = c.x + c.velX
		c.y = c.y + c.velY
		cords[i] = c
	}
	return cords
}

func size(cords []coord) int {
	minX, minY := 100000, 100000
	maxX, maxY := -100000, -100000
	for _, c := range cords {
		//fmt.Println(c.x)
		if minX > c.x {
			minX = c.x
		}
		if minY > c.y {
			minY = c.y
		}
		if maxX < c.x {
			maxX = c.x
		}
		if maxY < c.y {
			maxY = c.y
		}
	}
	return (maxX - minX) + (maxY - minY)
}

func pointInList(x, y int, cords []coord) bool {
	for _, c := range cords {
		if c.x == x && c.y == y {
			return true
		}
	}
	return false
}

func print(cords []coord) {
	minX, minY := 100000, 100000
	maxX, maxY := -100000, -100000

	for _, c := range cords {
		//fmt.Println(c.x)
		if minX > c.x {
			minX = c.x
		}
		if minY > c.y {
			minY = c.y
		}
		if maxX < c.x {
			maxX = c.x
		}
		if maxY < c.y {
			maxY = c.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if pointInList(x, y, cords) == true {
				fmt.Print(string("*"))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func main() {
	file, _ := os.Open("input.txt")
	scan := bufio.NewScanner(file)

	var cords []coord
	for scan.Scan() {
		line := scan.Text()
		var c coord
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d,  %d>", &c.x, &c.y, &c.velX, &c.velY)
		cords = append(cords, c)
	}
	i := 0
	minArea := 1000000
	minCord := make([]coord, len(cords))
	for y := 0; y < 100000; y++ {
		cords = tic(cords)
		temp := size(cords)
		if minArea > temp {
			copy(minCord, cords)
			minArea = temp
			i = y
		}
	}
	print(minCord)
	fmt.Println(i + 1)

}
