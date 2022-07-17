package sliceRotate

import "fmt"

// "-" means left, "+" means right
func SliceRotate(argSlice []int, rotation int) []int {
	rotatedSlice := make([]int, len(argSlice))
	if rotation > 0 {
		rotatedSlice := argSlice[len(argSlice)-rotation : len(argSlice)]
		rotatedSlice = append(rotatedSlice, argSlice...)
	} else {
		rotatedSlice := argSlice[len(argSlice)-rotation : len(argSlice)]
		rotatedSlice = append(rotatedSlice, argSlice...)
	}

	fmt.Println(rotatedSlice)

	return rotatedSlice[:len(argSlice)]
}
