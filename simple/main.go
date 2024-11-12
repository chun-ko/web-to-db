package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Example call
// curl "http://localhost:8080/distribute?address=123"
func main() {
	http.HandleFunc("/distribute", httpHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func httpHandler(response http.ResponseWriter, request *http.Request) {
	// Extract value from Request
	query := request.URL.Query()
	address := query.Get("address")

	// Connect to DB
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "docker")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("fail to open postgres db %v", err)
	}
	defer db.Close()

	// Write record
	insertStmt := fmt.Sprintf(`insert into "addresses"("name") values(%s)`, address)
	_, err = db.Exec(insertStmt)
	if err != nil {
		log.Fatalf("fail to insert record %v", err)
	}
	response.Write([]byte("ok"))
}
