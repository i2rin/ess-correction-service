package main

import (
	"ess-server/pkg/db"
	"time"
)

func main() {
	db.InitDB()
	time.Sleep(30 * time.Second)
}
