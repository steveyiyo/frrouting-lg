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
	// io.WriteString(w, "https://github.com/steveyiyo/frrouting-lg")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
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
			io.WriteString(w, "Running ping...\n\n")
			io.WriteString(w, ping(IP))
		}
		if action == "traceroute" {
			io.WriteString(w, "Running traceroute...\n\n")
			io.WriteString(w, traceroute(IP))
		}
		if action == "mtr" {
			io.WriteString(w, "Running mtr...\n\n")
			io.WriteString(w, mtr(IP))
		}
		if action == "bgpsummary" {
			io.WriteString(w, "Checking BGP Summary...\n\n")
			io.WriteString(w, bgpsummary())
		}
		if action == "routev4" {
			io.WriteString(w, "Checking BGP Route...\n\n")
			io.WriteString(w, routev4(IP))
		}
		if action == "routev6" {
			io.WriteString(w, "Checking BGP Route...\n\n")
			io.WriteString(w, routev6(IP))
		}
	}
}

func ping(IP string) string {
	fmt.Print("Running ping...\n")

	c := "ping -O -c 5 " + IP
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

	c := "mtr -G 2 -c 5 -erwbz " + IP
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

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("FRRouting Looking Glass\n")
	fmt.Print("Port listing at 32991\n")
	fmt.Print("Repo: https://github.com/steveyiyo/frrouting-lg\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://network.steveyi.net/looking-glass\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":32991", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
