package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("www/index.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createProfileHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("www/createProfile.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func addProfileContentHandler(writer http.ResponseWriter, request *http.Request) {
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	options := os.O_APPEND
	file, err := os.OpenFile("files/profile.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, fname, lname)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/", http.StatusFound)
}

func makeConnectionHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("www/connectionsPage.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)

}

func applyJobsHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("www/applyJobs.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/createProfile", createProfileHandler)
	http.HandleFunc("/addProfileContent", addProfileContentHandler)
	http.HandleFunc("/makeConnection", makeConnectionHandler)
	http.HandleFunc("/applyJobs", applyJobsHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
