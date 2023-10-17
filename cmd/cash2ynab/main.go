package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fileToConvert := "tests/utils/examples/ynab_report_one_transaction.csv"
	cmd := exec.Command("cat", fileToConvert)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running cat command, %v\n", err)
		log.Fatalf("output, %s\n", string(output))
		return
	}

	fmt.Print(string(output))
}
