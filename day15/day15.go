package main

import (
	"bufio"
	"os"
)

type cord struct {
	x, y int
}

type info struct {
	walkable bool
}

type animal struct {
	hp      int
	species string
	c       cord
	alive   bool
}

func animalNearby(animals []animal, a animal) (string, animal) {
	for _, an := range animals {
		if an.c.x == a.c.x && an.c.y == an.c.y+1 {
			return "up", an
		} else if an.c.x == a.c.x && an.c.y == an.c.y-1 {
			return "down", an
		} else if an.c.x == a.c.x-1 && an.c.y == an.c.y {
			return "left", an
		} else if an.c.x == a.c.x+1 && an.c.y == an.c.y {
			return "right", an
		}
	}
	return "none", a
}

func move(maps map[cord]info, animals []animal, a animal) ([]info, animal) {
	var c cord
	c = a.c
	c.y = c.y - 1
	if maps[c].walkable == true {
		a.c = c
		maps[c].walkable = false
		return maps, a
	}

}

func main() {
	file, _ := os.Open("input.txt")
	scan := bufio.NewScanner(file)
	maps := make(map[cord]info)
	var animals []animal
	x, y := 0, 0
	for scan.Scan() {
		line := scan.Text()
		for c := range line {
			var tempCord cord
			var m info
			var a animal
			tempCord.x = x
			tempCord.y = y
			if string(c) == "#" {
				m.c = tempCord
				m.walkable = false
			} else if string(c) == "." {
				m.c = tempCord
				m.walkable = true
			} else if string(c) == "G" {
				a.species = "goblin"
				a.hp = 200
				a.c = tempCord
				a.alive = true

				m.c = tempCord
				m.walkable = true
				animals = append(animals, a)
			} else if string(c) == "E" {
				a.species = "elf"
				a.hp = 200
				a.c = tempCord
				a.alive = true

				m.c = tempCord
				m.walkable = true
				animals = append(animals, a)
			}
			maps[tempCord] = m
			x++
		}
		y++
	}
}
