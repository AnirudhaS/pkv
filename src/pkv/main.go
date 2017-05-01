package main

import (
	"fmt"
)

func main() {
	migrate()
	set("key_1", "1")
	first_val, _ := get("key_1")
	fmt.Printf("key_1 => %v", first_val.Value)

	set("key_1", "2")
	second_val, _ := get("key_1")
	fmt.Printf("key_1 => %v", second_val.Value)

	set("key", "Cs")
	third_val, _ := get("key")
	fmt.Printf("third_val: %v", third_val.Value)

	val, _ := increment("key_1", 2)
	fmt.Printf("third_val: %v", val)
	fourth_val, _ := get("key_1")
	fmt.Printf("third_val: %v", fourth_val.Value)

	err := delete("key_1")
	fmt.Printf("%v", err)
}
