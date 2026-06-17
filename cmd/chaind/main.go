package main

import (
	"flag"
	"fmt"

	"github.com/mxmkiv/go-blockchain/internal/node"
)

func main() {
	/*

		TODO
		add config package


	*/

	launchMode := flag.String("mode", "", "launch mode --local or --global")
	flag.Parse()

	if *launchMode == "" {
		fmt.Println("no launch parameter (--local or --global), launch local")
		*launchMode = "local"
	}

	var ExtrenalIP string
	if *launchMode == "global" {
		fmt.Print("Enter your external IP address: ")
		fmt.Scan(&ExtrenalIP)
	}

	n := node.NewNode(*launchMode, 10, 64)
	fmt.Println("node created")
	n.Start()
	select {}

}
