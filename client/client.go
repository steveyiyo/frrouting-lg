package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var (
	SourceIP = "1.1.1.1"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`FRRouting Looking Glass\n`))

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

	} else {
		// 判斷
		action := r.Form["action"][0]
		IP := r.Form["IP"][0]
		if action == "ping" {
			ping(IP)
			// fmt.Fprintf(w, "Unknow")
		}
		if action == "traceroute" {
			traceroute(IP)
			fmt.Fprintf(w, "Unknow")
		}
		if action == "mtr" {
			mtr(IP)
			fmt.Fprintf(w, "Unknow")
		}
		if action == "bgpsummary" {
			bgpsummary()
			fmt.Fprintf(w, "Unknow")
		}
		if action == "routev4" {
			routev4(IP)
			fmt.Fprintf(w, "Unknow")
		}
		if action == "routev6" {
			routev6(IP)
			fmt.Fprintf(w, "Unknow")
		}
	}
}

func ping(IP string) (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "ping -I", SourceIP, "-O -c 10", IP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func traceroute(IP string) (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "traceroute", IP, "-a", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func mtr(IP string) (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "mtr -G 2 -c 10 -erwbz", IP, "--address", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func bgpsummary() (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp summary'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func routev4(IP string) (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show ip bgp $ip'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func routev6(IP string) (string, error) {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp $ip'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return string(out), err
	}
	fmt.Println(string(out))
	return string(out), err
}

func main() {
	http.HandleFunc("/", lg)

	err := http.ListenAndServe(":32991", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
