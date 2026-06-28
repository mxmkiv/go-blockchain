package main

import (
	"flag"
	"fmt"

	"github.com/mxmkiv/go-blockchain/internal/logger"
	"github.com/mxmkiv/go-blockchain/internal/miner"
	"github.com/mxmkiv/go-blockchain/internal/node"
	"github.com/mxmkiv/go-blockchain/internal/p2p"
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

	// services
	zapLogger := logger.NewLogger()
	msgManager := p2p.NewManager()
	miner := miner.New()

	n := node.NewNode(*launchMode, zapLogger, msgManager, miner, 10, 64)
	fmt.Println("node created")
	n.Start()
	select {}

}
