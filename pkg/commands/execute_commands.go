package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mattn/go-colorable"
)

var (
	Stdout = colorable.NewColorableStdout() // add a colorable std out
	Stderr = colorable.NewColorableStderr() // add a colorable std err
)

// colorizeLevel function for send (colored or common) message to output.
func colorizeLevel(level string) string {
	// Define variables.
	var (
		red         = "\033[0;31m"
		green       = "\033[0;32m"
		yellow      = "\033[1;33m"
		noColor     = "\033[0m"
		color, icon string
	)

	// Switch color.
	switch level {
	case "success":
		color = green
		icon = "[OK]"
	case "error":
		color = red
		icon = "[ERROR]"
	case "info":
		color = yellow
		icon = "[INFO]"
	default:
		color = noColor
	}

	// Send common or colored caption.
	return fmt.Sprintf("%s%s%s", color, icon, noColor)
}
func ShowError(text string) error {
	return fmt.Errorf("%s%s", colorizeLevel("error"), text)
}

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
		go func() {
			for scanner.Scan() {
				ShowMessage("", scanner.Text(), false, false)
			}
		}()
	}

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return ShowError(errWait.Error())
	}

	return nil
}

// ShowMessage function for showing output messages.
func ShowMessage(level, text string, startWithNewLine, endWithNewLine bool) {
	// Define variables.
	var startLine, endLine string

	if startWithNewLine {
		startLine = "\n" // set a new line
	}

	if endWithNewLine {
		endLine = "\n" // set a new line
	}

	// Formatting message.
	message := fmt.Sprintf("%s %s %s %s", startLine, colorizeLevel(level), text, endLine)

	// Return output.
	_, err := fmt.Fprintln(Stdout, message)
	if err != nil {
		return
	}
}
