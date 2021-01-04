package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "FRR Looking Glass\n\n")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

	} else {
		action := "NULL"
		IP := "NULL"
		for key, values := range r.Form {
			if key == "action" {
				action = values[0]
			} else {
				IP = values[0]
			}
		}
		fmt.Println("action: ", action)
		fmt.Println("IP: ", IP)
		if action == "ping" {
			io.WriteString(w, ping(IP))
			fmt.Fprintf(w, IP)
		}
		if action == "traceroute" {
			io.WriteString(w, traceroute(IP))
		}
		if action == "mtr" {
			io.WriteString(w, mtr(IP))
		}
		if action == "bgpsummary" {
			io.WriteString(w, bgpsummary())
		}
		if action == "routev4" {
			io.WriteString(w, routev4(IP))
		}
		if action == "routev6" {
			io.WriteString(w, routev6(IP))
		}
	}
}

func ping(IP string) string {
	fmt.Print("Running ping...\n")

	c := "ping -O -c 10 " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func traceroute(IP string) string {
	fmt.Print("Running traceroute...\n")

	c := "traceroute " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func mtr(IP string) string {
	fmt.Print("Running mtr...\n")

	c := "mtr -G 2 -c 10 -erwbz " + IP
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func bgpsummary() string {
	fmt.Print("Checking BGP Summary...\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp summary'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func routev4(IP string) string {
	fmt.Print("Checking BGP Route...\n")

	c := "vtysh -c 'show ip bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func routev6(IP string) string {
	fmt.Print("Checking BGP Route...\n")

	c := "vtysh -c 'show bgp " + IP + "'"
	cmd := exec.Command("bash", "-c", c)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func main() {
	http.HandleFunc("/", lg)

	err := http.ListenAndServe(":32991", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
