package main

import (
	"os"

	"github.com/evoaway/erc20-transfers-storage-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
