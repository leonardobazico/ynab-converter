// this is a integration test
// which run the cash2ynab cli command and check the output

package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCash2ynab(t *testing.T) {
	t.Parallel()

	projectFolderPath := getProjectFolderPath(t)
	cli := buildCli(t, projectFolderPath)
	t.Run("should run cash2ynab command and get the output", func(t *testing.T) {
		// Given
		cmd := exec.Command(cli, "tests/utils/examples/cash_app_report_one_transaction.csv")
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		assert.NoError(t, err)
		assert.Equal(t,
			"Date,Payee,Memo,Amount\n10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.9\n",
			string(output))
	})
}

func getProjectFolderPath(tb testing.TB) string {
	tb.Helper()

	testFolder, _ := os.Getwd()
	tb.Logf("Current test path: %s", testFolder)
	projectFolderPath := filepath.Dir(filepath.Dir(testFolder))
	tb.Logf("Current project folder: %s", projectFolderPath)

	return projectFolderPath
}

func buildCli(tb testing.TB, projectFolderPath string) string {
	tb.Helper()

	cliPath := "bin/cash2ynab-test"
	cmd := exec.Command("go", "build", "-o", cliPath, "cmd/cash2ynab/main.go")
	cmd.Dir = projectFolderPath
	err := cmd.Run()
	if err != nil {
		tb.Fatal(err)
	}

	return cliPath
}
