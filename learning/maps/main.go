package main

import "fmt"

// maps are essentially the same as javascript objects and python dicts
// except all the keys must share the same time, and all the values must share the same type
func main() {
	var dict1 map[string]int

	dict2 := map[string]int{
		"key_data_type": 5,
		"new":           21,
	}

	dict1["delete_this_key"] = 21
	delete(dict1, "delete_this_key")
	delete(dict2, "key_data_type")
}

func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// what's the difference between a map and a struct?
// All keys in a map must be the same type, they can be different in a struct
// In a map, all the keys are indexed, meaning we can iterate over all key value pairs, we cannot with structs
// A struct is a value type (pointers) whereas a map is a refernce type (no pointers)
// We should use a map to represent a collection of closely related properties, like colors. Mapping a color and it's hex code
