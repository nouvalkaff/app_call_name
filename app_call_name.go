package main

import (
	"fmt"

	go_call_name "github.com/nouvalkaff/mod_call_name/v2"
)

//go get github.com/nouvalkaff/mod_call_name/v2@v2.5.1
// use go mod tidy to keep it works

func main() {
	fmt.Println(go_call_name.CallName("Mona", 22, "Okay"))
}