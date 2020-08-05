package rangerepo

import (
	"fmt"
)

func CallRange() {
	var t []int = []int{1, 2, 3}

	for i, y := range t {
		a := i
		b := y
		fmt.Println(a, b, "rangetest")
	}
}
