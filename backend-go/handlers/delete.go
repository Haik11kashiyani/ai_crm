
package handlers

import (
	"database/sql"   // Needed to use *sql.DB and Exec
	//"fmt"            // Used to send response messages
	"log"            // Used for logging errors
	"net/http"       // To handle HTTP requests
)

func DeleteCustomer(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		query:="delete from customers where id=$1"
		_,err:=db.Exec(query,1)
		if err!=nil{
			log.Fatal("delete ma vandho 6 vadil ",err)
		}
	}
	
}