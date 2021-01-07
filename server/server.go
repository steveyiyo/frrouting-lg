package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type router []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func lg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// io.WriteString(w, "https://github.com/steveyiyo/frrouting-lg")

	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

		fmt.Println(r.URL.Path)
		p := ".." + r.URL.Path
		if p == "../" {
			p = "../static/v1.html"
		}
		if p == "../v1" {
			p = "../static/v1.html"
		}
		if p == "../v2" {
			p = "../static/v2.html"
		}
		http.ServeFile(w, r, p)

	} else if r.Method == "POST" {
		err := r.ParseMultipartForm(5 * 1024 * 1024 * 1024)
		if err != nil {
			fmt.Println("Error ParseMultipartForm: ", err)
			return
		}

		data, err := ioutil.ReadFile("router.json")

		var router router
		json.Unmarshal(data, &router)

		// fmt.Println(router)

		Router := "NULL"
		Action := "NULL"
		IP := "NULL"

		Router = r.FormValue("Router")
		Action = r.FormValue("Action")
		IP = r.FormValue("IP")

		fmt.Println("Action: ", Action)
		fmt.Println("IP: ", IP)
		fmt.Println("Router: ", Router)

		var RouterIP string
		RouterIP = "NULL"
		for n := range router {
			if router[n].Name == Router {
				fmt.Println("Router IP: ", router[n].IP)
				RouterIP = router[n].IP
			}
		}

		RouterURL := "http://" + RouterIP + ":32991"
		resp, err := http.PostForm(RouterURL,
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
	fmt.Print("Demo: https://lg.steveyi.net\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":32280", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
