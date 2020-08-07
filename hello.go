package main

import (
	"fmt"

	"rabie.com/hello/arrayRepo"
	"rabie.com/hello/forRepo"
	"rabie.com/hello/functionrepo"
	"rabie.com/hello/mapRepo"
	"rabie.com/hello/pointerRepo"
	"rabie.com/hello/pointerreceiverRepo"
	"rabie.com/hello/rangerepo"
	"rabie.com/hello/sliceRepo"
	"rabie.com/hello/structRepo"
)

func main() {
	fmt.Println("range //////////////////////////////////")
	rangerepo.CallRange()
	fmt.Println("slice //////////////////////////////////")
	sliceRepo.CallSlice()
	fmt.Println("slice of slice //////////////////////////////////")
	sliceRepo.CallSliceOfSlice()
	fmt.Println("pointer //////////////////////////////////")
	pointerRepo.CallPointer()
	fmt.Println("struct //////////////////////////////////")
	a := structRepo.CallStruct()
	a.Brake_pedal = 1
	a.Gas_pedal = 2
	a.Steering_well = 3
	fmt.Println("struct car value:", a)
	fmt.Println("for //////////////////////////////////")
	forRepo.CallFor()
	fmt.Println("map //////////////////////////////////")
	mapRepo.CallMap()
	fmt.Println("map and for /////////////////////////////////////////")
	forRepo.CallForTool(mapRepo.CallMapTool())
	fmt.Println("delete map /////////////////////////////////////////")
	mapRepo.CallDeleteMap("js")
	fmt.Println("sheck map /////////////////////////////////////////")
	mapRepo.CallSheckMap("c#")
	fmt.Println("make new map /////////////////////////////////////////")
	mapRepo.CallMakeMap("ts", 4)
	fmt.Println("call array /////////////////////////////////////////")
	arrayRepo.CallArray()
	fmt.Println("functions /////////////////////////////////////////")
	functionrepo.SimpleFunction()
	functionrepo.FunctionWhitParam(1, 2)
	functionrepo.ContinuedFunction(1, 2, 3)
	fmt.Printf("add function x+y= %v", functionrepo.Calladd(1, 2))
	functionrepo.Func1()
	functionrepo.Func3()
	functionrepo.CallfuncReturnFunc(2)()

	v := pointerreceiverRepo.Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

}

// type car struct {
// 	gas_pedal     int
// 	brake_pedal   int
// 	steering_well int
// }

// var y [4]int = [4]int{1, 2, 3, 4}
// var a = 1
// var b = &a // memory address

// func add(x int, y int) int {

// 	return x + y
// }

// func main() {

// 	var z []int = y[1:4]
// 	fmt.Println(z)
// 	fmt.Println(len(z), "len")
// 	fmt.Println(cap(z), "cap")
// 	fmt.Println(z[:cap(z)])
// 	fmt.Println(z[1:])
// 	car_a := car{55, 26, 31}
// 	var d = &car_a.brake_pedal

// 	fmt.Println(car_a.brake_pedal)
// 	*d = 100
// 	fmt.Println(car_a)
// 	fmt.Println(d)

// 	// fmt.Println(morestring.ReverseRunes("rabie"))
// 	// fmt.Println(cmp.Diff("rabie", 1.1234))
// 	// time.Sleep(2 * time.Second)

// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Print(rand.Intn(5))
// 	// fmt.Println()
// 	// fmt.Print(math.Pi)
// 	// fmt.Println()
// 	// a += 1
// 	// fmt.Println(a)
// 	// defer fmt.Println(add(1, 2), "+") //wait for all code to excecute than it work i used on boucle for it while excecute operation stacking
// 	// fmt.Println(math.Pow(3, 2))

// 	// fmt.Print("Go runs on ")
// 	// switch os := runtime.GOOS; os {
// 	// case "darwin":
// 	// 	fmt.Println("OS X.")
// 	// case "linux":
// 	// 	fmt.Println("Linux.")
// 	// default:
// 	// 	// freebsd, openbsd,
// 	// 	// plan9, windows...
// 	// 	fmt.Printf("%s.\n", os)
// 	// }

// 	// fmt.Println("When's Saturday?")
// 	// today := time.Now().Weekday()
// 	// switch time.Saturday {
// 	// case today + 0:
// 	// 	fmt.Println("Today.")
// 	// case today + 1:
// 	// 	fmt.Println("Tomorrow.")
// 	// case today + 2:
// 	// 	fmt.Println("In two days.")
// 	// default:
// 	// 	fmt.Println("Too far away.")
// 	// }
// 	// fmt.Println(time.Now().Weekday() + 1)
// 	// fmt.Println(time.Saturday)
// 	// fmt.Println(time.Now().Hour())
// 	fmt.Println(*b)
// 	*b = 5
// 	fmt.Println(a)
// 	fmt.Println(*b)
// 	a = 3

// 	fmt.Println(a)
// 	fmt.Println(*b)

// 	fmt.Println(b)
// }

// func main() {

// 	fmt.Print(rand.Intn(90), ",")
// 	fmt.Print(rand.Intn(100))
// 	fmt.Println()

// 	fmt.Println(rand.Float64())

// 	fmt.Print((rand.Float64()*5)+5, ",")
// 	fmt.Print((rand.Float64() * 5) + 5)
// 	fmt.Println()

// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)

// 	fmt.Print(r1.Intn(100), ",")
// 	fmt.Print(r1.Intn(100))
// 	fmt.Println()

// 	s2 := rand.NewSource(42)
// 	r2 := rand.New(s2)
// 	fmt.Print(r2.Intn(100), ",")
// 	fmt.Print(r2.Intn(100))
// 	fmt.Println()
// 	s3 := rand.NewSource(42)
// 	r3 := rand.New(s3)
// 	fmt.Print(r3.Intn(100), ",")
// 	fmt.Print(r3.Intn(100))
// }
