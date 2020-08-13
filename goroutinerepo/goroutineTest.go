package goroutinerepo

import (
	"fmt"
	"time"
)

// GoroutineTest : function that alow you to print i and wait for declared time
func GoroutineTest(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(s, i)
	}
}
