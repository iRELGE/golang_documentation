package mapRepo

import "fmt"

var a map[string]int = map[string]int{
	"c#": 1,
	"go": 2,
	"js": 3,
}

func CallMap() {

	fmt.Println(a)
	fmt.Println(a["go"])

}
func CallMapTool() (b map[string]int) {
	b = a
	return
}
func CallDeleteMap(key string) {
	delete(a, key)
	fmt.Printf("map %s has been delete it", key)
	fmt.Println()
	fmt.Println(a)
}
func CallSheckMap(key string) {
	val, ok := a[key]
	if ok == true {
		fmt.Printf("key %s exist", key)
		fmt.Println("key value:", val)

	} else {
		fmt.Printf("key %s doesn't exist", key)

	}
}
