package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

var (
	SourceIP = "1.1.1.1"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "FRRouting Looking Glass\n\n")

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

	} else {
		// 判斷
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
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "traceroute" {
			io.WriteString(w, traceroute(IP))
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "mtr" {
			io.WriteString(w, mtr(IP))
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "bgpsummary" {
			io.WriteString(w, bgpsummary())
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "routev4" {
			io.WriteString(w, routev4(IP))
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "routev6" {
			io.WriteString(w, routev6(IP))
			// fmt.Fprintf(w, "Unknow")
		}
	}
}

func ping(IP string) string {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "ping -I", SourceIP, "-O -c 10", IP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func traceroute(IP string) string {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "traceroute", IP, "-a", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func mtr(IP string) string {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "mtr -G 2 -c 10 -erwbz", IP, "--address", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func bgpsummary() string {
	fmt.Print("Client has been start!\n")

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
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show ip bgp $ip'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "An error occurred"
	}
	fmt.Println(string(out))
	return string(out)
}

func routev6(IP string) string {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp $ip'")
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
