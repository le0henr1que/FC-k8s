package main 

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		name := os.Getenv("NAME")
		age := os.Getenv("AGE")
		fmt.Fprintf(w, "Hello %s, you are %s years old", name, age)
		
	})

	http.ListenAndServe(":3000", nil)
}