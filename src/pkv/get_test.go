package main

import "testing"

func TestEmptyGet(t *testing.T) {
	_, err := get("ss")
	if err == nil {
		t.Error("Empty get() test failed.")
	}
}

func TestGet(t *testing.T) {
	db := connectDB()
	db.Create(&KeyValue{Key: "test", Value: "pass"})
	kv, err := get("test")
	if err != nil {
		t.Error("Expected ", "pass", "got Error.")
	}
	if kv.Value != "pass" {
		t.Error("Expected ", "pass", "got ", kv.Value)
	}
}
