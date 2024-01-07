package main

import "golang.org/x/tour/pic"

// Pic draws a representation of a chosen function
func Pic(dx, dy int) [][]uint8 {
	// Creating slices with length dy and dx
	slice := make([][]uint8, dy)

	// Loop through the dy slice to make each element a slice of dx
	for y := range slice {
		sliceOfDx := make([]uint8, dx)
		for x := range sliceOfDx {
			slice[y] = append(slice[y], uint8(y*x))
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
