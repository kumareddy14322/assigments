package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Move for the registration !") // write data to response
}

func regis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("regis.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("gender:", r.Form["gender"])
		fmt.Println("address:", r.Form["address"])
	}
}

func main() {
	http.HandleFunc("/", welcome) // setting router rule
	http.HandleFunc("/regis", regis)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}