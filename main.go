package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// type JobsApplied struct {
// 	Title          string
// 	ADate          string
// 	ResponseStatus string
// 	JobURL         string
// }

type JobAppliedDetails struct {
	AppliedDetails []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func splitJobFields(appliedJobs []string) [][]string {
	var splittedJobs [][]string
	for _, jobs := range appliedJobs {
		tempSplit := strings.Split(jobs, " ")
		splittedJobs = append(splittedJobs, tempSplit)
	}
	return splittedJobs
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

func viewAppliedJobsHandler(writer http.ResponseWriter, request *http.Request) {
	appliedJobs := getStrings("files/appliedJobs.txt")
	html, err := template.ParseFiles("www/viewJobs.html")
	check(err)

	jobDetails := JobAppliedDetails{
		AppliedDetails: appliedJobs,
	}
	if len(appliedJobs) != 0 {
		err = html.Execute(writer, jobDetails)
	} else {
		err = html.Execute(writer, nil)
		check(err)
	}

	// splittedJobs := splitJobFields(appliedJobs)

	// if len(splittedJobs) != 0 {
	// 	var jobsAppliedArray []JobsApplied
	// 	for _, jobs := range splittedJobs {
	// 		tempValue := JobsApplied{
	// 			Title:          jobs[0],
	// 			ADate:          jobs[1],
	// 			ResponseStatus: jobs[2],
	// 			JobURL:         jobs[3],
	// 		}
	// 		jobsAppliedArray = append(jobsAppliedArray, tempValue)
	// 	}
	// 	err = html.Execute(writer, jobsAppliedArray)
	// 	check(err)
	// } else {
	// 	err = html.Execute(writer, nil)
	// 	check(err)
	// }

}

func addAppliedJobsFormHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("www/addAppliedJobsForm.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func addAppliedJobsHandler(writer http.ResponseWriter, request *http.Request) {
	jTitle := request.FormValue("jtitle")
	aDate := request.FormValue("adate")
	status := request.FormValue("status")
	jURL := request.FormValue("jurl")
	options := os.O_APPEND
	file, err := os.OpenFile("files/appliedJobs.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, jTitle, aDate, status, jURL)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/", http.StatusFound)

}

func main() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/createProfile", createProfileHandler)
	http.HandleFunc("/addProfileContent", addProfileContentHandler)
	http.HandleFunc("/makeConnection", makeConnectionHandler)
	http.HandleFunc("/applyJobs", applyJobsHandler)
	http.HandleFunc("/viewAppliedjobs", viewAppliedJobsHandler)
	http.HandleFunc("/addAppliedJobsForm", addAppliedJobsFormHandler)
	http.HandleFunc("/addAppliedJobs", addAppliedJobsHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
