package main

import (
	"MyProject/api"
	"database/sql"
	"log"
	"net/http"
)

func main() {

	dsn := "root:mysql@tcp(localhost)MySQl80?parseTime=true)"
	db, err := sql.Open("msql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(db)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
