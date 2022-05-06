package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// ExecCommand function to execute a given command.
func ExecCommand(command string, options []string, silentMode bool) error {
	// Checking for nil.
	if command == "" || options == nil {
		return fmt.Errorf("no command to execute")
	}

	// Create buffer for stderr.
	stderr := &bytes.Buffer{}

	// Collect command line.
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd.
	cmd.Stderr = stderr

	// Create a new reader.
	cmdReader, errStdoutPipe := cmd.StdoutPipe()
	if errStdoutPipe != nil {
		return ShowError(errStdoutPipe.Error())
	}

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return ShowError(errStart.Error())
	}

	// Create a new scanner and run goroutine func with output, if not in silent mode.
	if !silentMode {
		scanner := bufio.NewScanner(cmdReader)
		// go func() {
		for scanner.Scan() {
			ShowMessage("", scanner.Text(), false, false)
		}
		// }()
	}

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return ShowError(errWait.Error())
	}
	return nil
}

func ExecCommandWithOutput(command string, options []string) (string, error) {
	cmd := exec.Command(command, options...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Printf("Out:\n%s, err:\n%s", cmd.Stdout, cmd.Stderr)
	if err != nil {
		fmt.Println("errror", err)
		return "", nil
	}
	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, nil
}
