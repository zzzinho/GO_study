package main

import (
	_ "database/sql"

	"github.com/study/api_structure/db"
	"github.com/study/api_structure/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
