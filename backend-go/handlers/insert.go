// File: handlers/insert.go

package handlers

import (
	"database/sql"   // Needed to use *sql.DB and Exec
	"fmt"            // Used to send response messages
	"log"            // Used for logging errors
	"net/http"       // To handle HTTP requests

	
)

// InsertCustomer handles inserting a customer into the database
func InsertCustomer(db *sql.DB) http.HandlerFunc {
	// Return an actual HTTP handler function
	return func(w http.ResponseWriter, r *http.Request) {
		
		// Static values for now
		nameOfUser := "hardik"
		email := "hardik123@gmail.com"
		mobile := "1234567890"

		// SQL query with placeholders ($1, $2, $3) to prevent SQL injection
		query := "INSERT INTO customers(name, email, mobile) VALUES($1, $2, $3)"

		// Execute the query using db.Exec
		_, err := db.Exec(query, nameOfUser, email, mobile)
		if err != nil {
			log.Fatal("❌ Error inserting into customers table:", err)
		}

		// Send response back to client
		fmt.Fprintln(w, "✅ Inserted customer successfully")
	}
}
