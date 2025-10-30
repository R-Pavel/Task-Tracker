package main

import (
	"fmt"
	"time"
)

const (
	ColorReset = "\033[0m"
	ColorRed = "\033[31m"
	ColorGreen = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue = "\033[34m"
	ColorGray = "\033[37m"
)

type Task struct {
	ID int `json: "id"`
	Description string `json: "description"`
	Status string `json: "status"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "UpdatedAt"`
}

