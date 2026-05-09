package main

import "fmt"

func main() {

	// map[keytype]valuetype{key:value, ...}
	var m1 = map[string]string{"first_name":"John"}

	fmt.Println(m1["first_name"])

	// make - nil - to create a empty map 
	var m2 = make(map[string]int)

	m2["apple"] = 10
	m2["banana"] = 20

	fmt.Println(m2)

	// delete(map_name, key) to delete a key and a value
	delete(m2, "apple")

	fmt.Println(m2)

	// can take value and bool check
	value, check := m2["banan"]

	fmt.Println(value, check)

}
