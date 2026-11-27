package main

import "fmt"

/* 
	golang supports mutiple returns 
	IMP: go functions are first class citizen functions 
	whihc means we can assign that to a variable and pass as a agrument also 
*/


/* Example 1 */
func add(a int, b int) int {
	return a + b
}

/* Example 2 */ 
func getLanguages() (string, string, string){
	return "golang", "javascript", "c"
}

/* Example 3 */
func processIt(fn func(a int) int){
	fn(1)
}

/* Example 4 */
func process() func(a int) int{
	return func(a int) int {
		return 5
	}
}

func main() {
	result := add(1, 5)

	fmt.Println(result)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int{
		return a
	}

	processIt(fn)

	newFn := process()
	val := newFn(5)

	fmt.Println(val)
}