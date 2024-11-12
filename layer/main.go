package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/distribute", httpHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func httpHandler(response http.ResponseWriter, request *http.Request) {
	// Extract value from Request
	query := request.URL.Query()
	address := query.Get("address")
	err := writeAddress(address)
	if err != nil {
		log.Fatal("Fail to write into db")
	}
	response.Write([]byte("ok"))
}

func writeAddress(address string) error {
	// Connect to DB
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "docker")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Write record
	insertStmt := fmt.Sprintf(`insert into "addresses"("name") values(%s)`, address)
	_, err = db.Exec(insertStmt)
	if err != nil {
		return err
	}
	return nil
}
