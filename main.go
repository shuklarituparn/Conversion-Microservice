//package main
//
//import (
//	"log"
//	"net/http"
//)
//
//func main() {
//	fs := http.FileServer(http.Dir("static"))
//	http.Handle("/", fs)
//	if err := http.ListenAndServe(":8085", nil); err != nil {
//		log.Println(err)
//	}
//}
//TO CHECK THE OS.EXEC COMMANDS

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
