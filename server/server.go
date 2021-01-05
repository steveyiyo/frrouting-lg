package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// io.WriteString(w, "https://github.com/steveyiyo/frrouting-lg")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		Router := "NULL"
		Action := "NULL"
		IP := "NULL"
		for key, values := range r.Form {
			if key == "Router" {
				Router = values[0]
			}
			if key == "IP" {
				IP = values[0]
			}
			if key == "Action" {
				Action = values[0]
			}
		}

		fmt.Println("Action: ", Action)
		fmt.Println("IP: ", IP)
		fmt.Println("Router: ", Router)

		resp, err := http.PostForm("http://10.121.211.254:32991",
			url.Values{"Action": {Action}, "IP": {IP}})
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		io.WriteString(w, string(body))
	}
}

func main() {
	http.HandleFunc("/", lg)

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("FRRouting Looking Glass\n")
	fmt.Print("Port listing at 32280\n")
	fmt.Print("Repo: https://github.com/steveyiyo/frrouting-lg\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://network.steveyi.net/looking-glass\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":32280", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
