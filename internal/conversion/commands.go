package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("mkdir", "hello")
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}

//The first value is the utility we use and the next is the arguments we pass
//In a single directory can't have mutliple packages
