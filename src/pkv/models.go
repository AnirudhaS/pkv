package main

import "time"

// Basic key value storage.
type KeyValue struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Key       string
	Value     string
	TTL       time.Time
	TTLSet    bool
}
