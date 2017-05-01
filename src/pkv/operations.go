package main

import (
	"errors"
	"fmt"
	"strconv"
)

func create(key string, value string) {
	db := migrate()
	db.Create(&KeyValue{Key: key, Value: value})
}

func get(key string) (*KeyValue, error) {
	db := connectDB()
	kv := &KeyValue{}
	db.Where("key = ?", key).First(kv)
	if kv.Key == "" {
		return nil, errors.New("key doesn't exist")
	}
	return kv, nil
}

func set(key string, value string) {
	db := connectDB()
	tx := db.Begin()
	defer tx.Commit()
	if kv, err := get(key); err != nil {
		tx.Create(&KeyValue{Key: key, Value: value})
	} else {
		kv.Value = value
		tx.Save(kv)
		// tx.Model(kv).Update("Value", value)
	}
}

func delete(key string) error {
	db := connectDB()
	tx := db.Begin()
	kv, err := get(key)
	if err != nil {
		return errors.New("No such key")
	}
	tx.Delete(KeyValue{}, "key = ?", kv.Key)
	tx.Commit()
	return nil
}

func exists(key string) bool {
	db := connectDB()
	kv := &KeyValue{Key: key}
	return db.NewRecord(kv)
}

func increment(key string, by int) (int, error) {
	db := connectDB()
	tx := db.Begin()
	defer tx.Commit()
	kv, err := get(key)
	if err != nil {
		return 0, fmt.Errorf("'%s' key does not exist", key)
	}
	value, err := strconv.Atoi(kv.Value)
	if err != nil {
		return 0, errors.New("Value is not a string")
	}
	set(key, strconv.Itoa(value+by))
	return value + by, nil
}
