package main

import "fmt"

// map -> hash, object
func main() {
	// // creating map string type with string key 
	// m := make(map[string]string)
	
	// // setting an element 
	// m["name"]="golang"
	// m["area"]="backend"
	
	// // get an element 
	// // fmt.Println(m["name"], m["area"])
	// // IMP: if any keys does not exist in the map then it returns zero value 
	
	// fmt.Println(m["mobile"]) -> empty string 
	
	// // creating map int type with string key 
	// m := make(map[string]int)
	// m["age"] = 30 
	// m["price"] = 50

	// fmt.Println(m["age"], m["price"])

	// fmt.Println("Before delete", m)
	// // delete 
	// delete(m, "price")

	// fmt.Println("After delete",m)

	// // clear a map 
	// fmt.Println(m)

	// without make function 

	m := map[string]float32{
		"price":15.00,
		"quantity":25.00,
		"total":25.5555,
	}

	fmt.Println(m)

	// check weather key exits in the map or not 
	v, is_key := m["price"]

	if is_key {
		fmt.Println("Key found", v)
	} else {
		fmt.Println("Key not found", v)
	}
}