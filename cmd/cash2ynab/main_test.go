// this is a integration test
// which run the cash2ynab cli command and check the output

package main_test

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCash2ynab(t *testing.T) {
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

func getProjectFolderPath(t testing.TB) string {
	_, testFilePath, _, _ := runtime.Caller(0)
	projectFolderPath := filepath.Dir(filepath.Dir(filepath.Dir(testFilePath)))
	t.Logf("Current test basepath: %s", projectFolderPath)
	return projectFolderPath
}

func buildCli(t testing.TB, projectFolderPath string) string {
	cliPath := "bin/cash2ynab-test"
	cmd := exec.Command("go", "build", "-o", cliPath, "cmd/cash2ynab/main.go")
	cmd.Dir = projectFolderPath
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return cliPath
}
