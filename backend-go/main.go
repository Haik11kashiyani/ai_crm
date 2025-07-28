package main

import (
	"database/sql"
	"fmt" //working like we can say this is for print statements
	"log" //helping to give error messages
	"net/http" //for handling HTTP requests or making the server
	"os" //for reading environment variables

	"github.com/joho/godotenv" //for loading environment variables from .env file
	_"github.com/lib/pq" // for to connect go language  with our postgresql db insort its a driver _ means we are only import it for run init()
	
	"backend-go/handlers"
)

func main(){
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	// geting the .env files filed and short in vairables like db usename and pass etc.

	db_uname := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	 coonstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable" , db_uname ,db_pass ,db_host ,db_port ,db_name)

	
	// we create two variable same time if got in constring error that will store in errorr_constring othervise it gives the resonse so after that we stores in db for forther oprations
	 
	 db, errorr_constring := sql.Open("postgres", coonstring)

	if errorr_constring != nil{
		log.Fatal("problem in connection " , errorr_constring)
	}

	// give an hello messege lets say to db if we got the respond that means db is connected.
	errmess := db.Ping()
	if errmess!=nil	{
		log.Fatal("db is not responed not connected with the db " , errmess)
	}
	
	http.HandleFunc("/insert", handlers.InsertCustomer(db))
	
	fmt.Println("✅ Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}