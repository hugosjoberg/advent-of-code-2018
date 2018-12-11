package main

import (
	"fmt"
	"strconv"
)

type coord struct {
	x, y, powerLevel int
}

func calculate(x, y, serialNumber int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel = powerLevel + serialNumber
	powerLevel = powerLevel * rackID
	temp := strconv.Itoa(powerLevel)
	if len(temp) < 2 {
		powerLevel = 0
		return powerLevel
	} else {
		powerLevel, _ = strconv.Atoi(string(temp[len(temp)-3]))
		powerLevel = powerLevel - 5
	}
	return powerLevel
}

func bestCoord(memory map[coord]int, gridSize int) (coord, int) {
	var bestCord coord
	totalPower := -10000
	for x := 0; x < 300-gridSize; x++ {
		for y := 0; y < 300-gridSize; y++ {
			tempPower := 0
			var tempCoord coord
			for tempX := x; tempX < x+gridSize; tempX++ {
				for tempY := y; tempY < y+gridSize; tempY++ {
					var c coord
					c.x = tempX
					c.y = tempY
					tempPower = tempPower + memory[c]
					if tempX == x && tempY == y {
						tempCoord = c
					}
				}
			}
			if tempPower > totalPower {
				totalPower = tempPower
				bestCord = tempCoord
			}
		}
	}
	return bestCord, totalPower
}

func main() {
	gridSerialNumber := 9110
	memory := make(map[coord]int)
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			var c coord
			c.x = x
			c.y = y
			powerLevel := calculate(x, y, gridSerialNumber)
			memory[c] = powerLevel
		}
	}

	tempCoord, _ := bestCoord(memory, 3)
	fmt.Println("Part 1:")
	fmt.Println(tempCoord)

	var bestCord coord
	bestPower := 0
	for i := 0; i <= 300; i++ {
		tempCoord, tempPower := bestCoord(memory, i)
		if tempPower > bestPower {
			bestPower = tempPower
			bestCord = tempCoord
			fmt.Println(i)
			fmt.Println(bestPower)
			fmt.Println(bestCord)
		}

	}
	fmt.Println("Part 2:")
	fmt.Println(bestPower)
	fmt.Println(bestCord)

}
