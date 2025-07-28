
package handlers

import (
	"database/sql"   // Needed to use *sql.DB and Exec
	"fmt"            // Used to send response messages
	"log"            // Used for logging errors
	"net/http"       // To handle HTTP requests
)

func SelectCustomer(db *sql.DB) http.HandlerFunc{

	return func(w http.ResponseWriter, r*http.Request){
		query:="select id,name from customers"
		row, err:=db.Query(query)
		if err!=nil{
			log.Fatal("select query ma vandho 6 vadil ",err)
		}

		 defer row.Close()
		 var result string
		 for row.Next(){

			 var id int 
			var name string
			err:=row.Scan(&id,&name)
			if err!=nil{
				log.Fatal("scan ma vandho 6 vadil ",err)

			}
			result+=fmt.Sprintf("id: %d, name: %s\n", id, name)
		}
		fmt.Fprintln(w,result)
	}
	
}