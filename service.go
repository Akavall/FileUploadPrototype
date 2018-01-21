package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("file_upload_prototype.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	http.HandleFunc("/", index)

	ip_address := "0.0.0.0:8080"
	fmt.Printf("Running on: %s\n", ip_address)

	err := http.ListenAndServe(ip_address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
