package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/valyala/fasthttp"
)

var (
	w        http.ResponseWriter
	r        *http.Request
	ctx      *fasthttp.RequestCtx
	SourceIP = "1.1.1.1"
	// addr     = flag.String("addr", ":8080", "TCP address to listen to")
	// compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	router := string(ctx.FormValue("router"))
	action := string(ctx.FormValue("action"))
	IP := string(ctx.FormValue("IP"))
	fmt.Println(router, action, IP)
}

func ping() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "ping -I", SourceIP, "-O -c 10", IP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func traceroute() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "traceroute", IP, "-a", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func mtr() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "mtr -G 2 -c 10 -erwbz", IP, "--address", SourceIP)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func bgpsummary() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp summary'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func routev4() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show ip bgp $ip'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func routev6() {
	fmt.Print("Client has been start!\n")

	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp $ip'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
