package main

import (
	"fmt"
	"os/exec"
)

func main() {

	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
