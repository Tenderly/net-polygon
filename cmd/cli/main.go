package main

import (
	"os"

	"github.com/tenderly/net-polygon/internal/cli"
	"github.com/tenderly/net-polygon/params"
)

func main() {
	params.UpdateBorInfo()
	os.Exit(cli.Run(os.Args[1:]))
}
