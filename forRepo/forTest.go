package forRepo

import "fmt"

func CallFor() {
	var a []int = []int{1, 4, 2, 3, 4}

	for _, y := range a {
		for _, y2 := range a {
			if y == y2 {
				fmt.Println(y, "y")

			}
		}
	}
}
func CallForTool(a map[string]int) {
	for i, j := range a {
		println("index:", i, "value:", j)
	}
}
