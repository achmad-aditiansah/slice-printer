package printer

import "fmt"

func PrintSlice(slice []int) {
	for _, v := range slice {
		fmt.Print(v, "\t")
	}
}

func PrintSlice2D(slice [][]int) {
	for i := range slice {
		for j := range slice[i] {
			fmt.Print(slice[i][j], "\t")
		}
	}
}