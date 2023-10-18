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
	outputString := string(output)
	if err != nil {
		log.Fatalf(
			"Error running cat command, %v\n"+
				"output, %s\n",
			err,
			outputString,
		)
	}

	fmt.Print(outputString)
}
