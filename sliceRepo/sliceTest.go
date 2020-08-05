package sliceRepo

import (
	"fmt"
	"time"
)

var a [4]string = [4]string{"rabia", "chaima", "chadi", "sara"}

func CallSlice() {
	b := a[0:2]
	fmt.Println("cap of b: ", cap(b), "len of b", len(b))
	fmt.Println("value of b:", b)
	c := b[1:4]
	fmt.Println("c table", c)
	c[0] = "change a"
	fmt.Println("change in source c reflex in a", a)
	b = append(b, "newBrother")
	fmt.Println("all value of b:", b)
	fmt.Println("all value of first table a", a)
	fmt.Println("whait for two second to see your new slice")
	time.Sleep(2 * time.Second)
	e := make([]int, 5, 5)
	fmt.Println("len of new slice e:", len(e), "cap of new slice e:", cap(e))
	fmt.Println("e values:", e)

}
func CallSliceOfSlice() {
	b := [][]string{
		[]string{"1", "*", "*"},
		[]string{"2", "*", "*"},
		[]string{"3", "*", "*"},
		[]string{"4", "*", "*"},
	}
	fmt.Println("normale value of slice b:", b)
	c := b[0:2]
	fmt.Println("slice c:", c)
}
