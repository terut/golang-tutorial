package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	parts := make([][]uint8, dy)
	for y := range parts {
		parts[y] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			parts[y][x] = uint8((x * y) / 2)
		}
	}
	return parts
}

func main() {
	pic.Show(Pic)
}
