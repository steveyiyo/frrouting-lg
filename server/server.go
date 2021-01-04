package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "FRR Looking Glass\n\n")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

	} else {
		router := "NULL"
		for key, values := range r.Form {
			if key == "router" {
				router = values[0]
			}
		}
		fmt.Println("router: ", router)
		action := "NULL"
		IP := "NULL"
		for key, values := range r.Form {
			if key == "action" {
				action = values[0]
			} else {
				IP = values[0]
			}
		}

		jsonFile, err := os.Open("router.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		resp, err := http.PostForm("http://", json.IP,
			url.Values{"action": {action}, "IP": {IP}})
		if err != nil {
			// handle error
		}
	}
}

func main() {
	http.HandleFunc("/", lg)

	err := http.ListenAndServe(":32280", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Json

}
