package functionrepo

import "fmt"

//SimpleFunction :simple function
func SimpleFunction() {
	fmt.Println("this is a simple function")
}

//FunctionWhitParam :function whit parameters
func FunctionWhitParam(x int, y int) {
	fmt.Printf("this is a function whit parameters x: %v and y: %v", x, y)
	fmt.Println()
}

//ContinuedFunction :continued function
func ContinuedFunction(a, b, c int) {
	fmt.Printf("this is a continued function whit paral a:%v , b: %v and c:%v ", a, b, c)
	fmt.Println()
}

//Func1 :function inside a function  call it directly and give it a param
func Func1() {
	fmt.Println("hi function 1")
	Second := func(x int) int {

		return x << x
	}(10)
	fmt.Println(Second)
}

/*Func3 :second variable take 7 whish is a value of param test then function second return a value -7 to func2 that take func as param*/
func Func3() {

	Second := func(x int) int {

		return -1 * x
	}
	func2(Second)

}

//function take function as parame
func func2(test func(int) int) {

	fmt.Println(test(7))
}

func test1(i int) int {
	return i * -1
}

//function take function as parame
func test2(test int) {

	fmt.Println(test1(9))
}

//Calladd :exrc add function
func Calladd(x int, y int) int {

	return x + y
}

//CallfuncReturnFunc : its a function that return a function
func CallfuncReturnFunc(x int) func() {
	return func() { fmt.Println("function has been retirned whit value x: ", x) }
}
