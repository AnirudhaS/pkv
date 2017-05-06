package main

import (
	"errors"
	"time"
)

type Hash struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	KeyValues []KeyValue
	CreatedAt time.Time
	UpdatedAt time.Time
}

func hashCreate(name string) (*Hash, error) {
	db := migrate()
	hash := Hash{Name: name}
	db.Create(&hash)
	return &hash, nil
}

func hashFind(name string) (*Hash, error) {
	db := connectDB()
	hash := &Hash{}
	db.Where("name = ?", name).First(hash)
	if hash.Name == "" {
		return nil, errors.New("No such hash")
	}
	return hash, nil
}

func hashFindOrCreate(name string) *Hash {
	hash, err := hashFind(name)
	if err != nil {
		hash, _ := hashCreate(name)
		return hash
	}
	return hash
}

func hashGet(hash *Hash, key string) (*KeyValue, error) {
	db := connectDB()
	kv := &KeyValue{}
	db.Model(hash).First(kv)
	if kv.Key != "" {
		return kv, nil
	}
	return nil, errors.New("Key not found")
}

func hashSet(name string, key string, value string) (*Hash, error) {
	hash := hashFindOrCreate(name)
	db := connectDB()
	// if err != nil {
	// 	hash, _ := hashCreate(name)
	// }
	kv, err := hashGet(hash, key)
	if err == nil {
		kv.Value = value
		db.Save(kv)
	} else {
		kv := &KeyValue{Key: key, Value: value, HashID: hash.ID}
		db.Create(kv)
	}
	return hash, nil
}
