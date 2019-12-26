package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	pic := make([][]uint8, dy)

	for y := range pic {
	
		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			// You can try (x+y)/2, x+y for more interesting 
			pic [y][x] =  uint8(x^y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
