package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Print("Client has been start!\n")
	cmd := exec.Command("bash", "-c", "vtysh -c 'show bgp summary'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
