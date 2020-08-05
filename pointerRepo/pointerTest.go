package pointerRepo

import "fmt"

var a int = 0

func CallPointer() {
	b := &a
	fmt.Println("memory adress that value will be stored", b)
	fmt.Println("value:", *b)
	*b = 3
	fmt.Println("new adress of a:", a)

}
