package main

import (
	"database/sql"
	"fmt" //working like we can say this is for print statements
	"log" //helping to give error messages
	"net/http" //for handling HTTP requests or making the server
	"os" //for reading environment variables

	"github.com/joho/godotenv" //for loading environment variables from .env file
	_"github.com/lib/pq" // for to connect go language  with our postgresql db insort its a driver _ means we are only import it for run init()
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
	
	http.HandleFunc("/add-customer", func(w http.ResponseWriter , r *http.Request) {
		
			name_of_user:="hardik"
			email:="hardik123@gmail.com"
			mobile:="1234567890"
		
		query:="insert into customers(name,email,mobile) values($1,$2,$3)"
		_,err:=db.Exec(query,name_of_user,email,mobile)
		if err!=nil{
			log.Fatal("problem in query in insert mota bhai",err)
		}
		fmt.Fprintln(w,"hellow hardik how are you? from db the db is connected suceesfullys") //this will write the response to the client

	})

	http.HandleFunc("/update",func(w http.ResponseWriter , r *http.Request){

		query:="update customers set name=$1 where id=$2 "
		newname:="hardik bhai kashiyani"
		_,err=db.Exec(query,newname, 1)
		if err !=nil{
			log.Fatal("update ma vandho 6 vadil ",err)
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request){
		query:="delete from customers where id=$1"
		_,err:=db.Exec(query,1)
		if err!=nil{
			log.Fatal("delete ma vandho 6 vadil ",err)
		}
	})
	fmt.Println("✅ Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}