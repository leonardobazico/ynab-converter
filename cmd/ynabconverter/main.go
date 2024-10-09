package main

import (
	"flag"
)

func main() {
	cashAppCmd := flag.NewFlagSet("cashapp", flag.ExitOnError)
	cashAppFile := cashAppCmd.String("file", "", "(required) path to CashApp csv file")

	cashApp(cashAppCmd, cashAppFile)
}
