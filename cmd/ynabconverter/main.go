package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cashAppCmd := flag.NewFlagSet("cashapp", flag.ExitOnError)
	cashAppFile := cashAppCmd.String("file", "", "(required) path to CashApp csv file")

	command := getCommandString()

	switch command {
	case "cashapp":
		cashApp(cashAppCmd, cashAppFile)
	default:
		fmt.Fprintf(os.Stderr, "Command available: cashapp")
		os.Exit(1)
	}
}

func getCommandString() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	return ""
}
