// this is a integration test
// which run the ynabconverter cli command and check the output

package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCash2ynab(t *testing.T) {
	t.Parallel()

	if os.Getenv("SKIP_INTEGRATION") == "true" {
		t.Skip("Skipping integration test")
	}

	projectFolderPath := getProjectFolderPath(t)
	cli := buildCli(t, projectFolderPath)
	t.Run("should run ynabconverter cashapp command and get the output", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(cli, "cashapp", "-file", "tests/utils/examples/cash_app_report_sample.csv")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.NoError(t, err)
		assert.Equal(t,
			"Date,Payee,Memo,Amount\n"+
				"10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90\n"+
				"01/10/2023,Some business name,PAYMENT SENT,-10.00\n",
			string(output))
	})

	t.Run("should fail when file does not exist", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(cli, "cashapp", "-file", "tests/utils/examples/does-not-exist.csv")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.Error(t, err)
		assert.Contains(
			t,
			string(output),
			"fail to open file: open tests/utils/examples/does-not-exist.csv: no such file or directory",
		)
	})

	t.Run("should fail when file is not provided", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(cli, "cashapp", "-file", "")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.Error(t, err)
		assert.Contains(
			t,
			string(output),
			"it is required to set a file to be converted",
		)
	})

	t.Run("should handle absolute path", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(cli, "cashapp", "-file", projectFolderPath+"/tests/utils/examples/cash_app_report_sample.csv")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.NoError(t, err)
		assert.Equal(t,
			"Date,Payee,Memo,Amount\n"+
				"10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90\n"+
				"01/10/2023,Some business name,PAYMENT SENT,-10.00\n",
			string(output))
	})

	t.Run("should handle relative path", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(cli, "cashapp", "-file", "./tests/utils/examples/cash_app_report_sample.csv")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.NoError(t, err)
		assert.Equal(t,
			"Date,Payee,Memo,Amount\n"+
				"10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90\n"+
				"01/10/2023,Some business name,PAYMENT SENT,-10.00\n",
			string(output))
	})

	t.Run("should handle level up relative path", func(t *testing.T) {
		t.Parallel()

		// Given
		cmd := exec.Command(
			cli,
			"cashapp",
			"-file",
			"../"+filepath.Base(projectFolderPath)+
				"/tests/utils/examples/cash_app_report_sample.csv",
		)
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.NoError(t, err)
		assert.Equal(t,
			"Date,Payee,Memo,Amount\n"+
				"10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90\n"+
				"01/10/2023,Some business name,PAYMENT SENT,-10.00\n",
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

	cliPath := "bin/ynabconverter-test"
	var cmd *exec.Cmd
	if _, present := os.LookupEnv("GOCOVERDIR"); present {
		cmd = exec.Command("go", "build", "-cover", "-o", cliPath, "./cmd/ynabconverter/...")
	} else {
		cmd = exec.Command("go", "build", "-o", cliPath, "./cmd/ynabconverter/...")
	}
	cmd.Dir = projectFolderPath
	err := cmd.Run()
	if err != nil {
		tb.Fatal(err)
	}

	return cliPath
}
