package main

import (
	"fmt" //working like we can say this is for print statements
	"log" //helping to give error messages
	"net/http" //for handling HTTP requests or making the server
	//"os" //for reading environment variables

	"github.com/joho/godotenv" //for loading environment variables from .env file
)

func main(){
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter , r *http.Request) {
		fmt.Fprintln(w,"hellow user how are you?") //this will write the response to the client
	})

	fmt.Println("✅ Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}