
package handlers

import (
	"database/sql"   // Needed to use *sql.DB and Exec
	//"fmt"            // Used to send response messages
	"log"            // Used for logging errors
	"net/http"       // To handle HTTP requests
)

func UpdateCustomer(db *sql.DB) http.HandlerFunc{


	return func(w http.ResponseWriter , r *http.Request){
		
		query:="update customers set name=$1 where id=$2 "
		newname:="hardik bhai kashiyani"

		_, err :=db.Exec(query,newname, 1)
		
		if err !=nil{
			log.Fatal("update ma vandho 6 vadil ",err)
		}
	}
}
