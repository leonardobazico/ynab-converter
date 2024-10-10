package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	t.Parallel()

	if os.Getenv("SKIP_INTEGRATION") == "true" {
		t.Skip("Skipping integration test")
	}

	projectFolderPath := getProjectFolderPath(t)
	cli := buildCli(t, projectFolderPath)

	t.Run("should show available commands when nothings is passed", func(t *testing.T) {
		// Given
		cmd := exec.Command(cli, "")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.Error(t, err)
		assert.Equal(t, "Command available: cashapp", string(output))
	})

	t.Run("should show available commands when command is not recognize", func(t *testing.T) {
		// Given
		cmd := exec.Command(cli, "do-nothing")
		cmd.Env = append([]string{}, os.Environ()...)
		cmd.Dir = projectFolderPath
		// When
		output, err := cmd.CombinedOutput()
		// Then
		require.Error(t, err)
		assert.Equal(t, "Command available: cashapp", string(output))
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
