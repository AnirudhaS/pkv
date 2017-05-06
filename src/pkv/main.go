package main

func main() {
	migrate()
	set("key", "first_value")
	set("key", "second_value")
	hashSet("first_hash", "first_key", "value_2")
	// hashCreate("first")
	// hashSet("first", "key", "value")
}
