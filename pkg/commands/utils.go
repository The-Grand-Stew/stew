package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mattn/go-colorable"
)

var (
	Stdout = colorable.NewColorableStdout() // add a colorable std out
	Stderr = colorable.NewColorableStderr() // add a colorable std err
)

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

// ShowError function for send error message to output.
func ShowError(text string) error {
	return fmt.Errorf("%s%s", colorizeLevel("error"), text)
}

// CalculateDurationTime func to calculate duration time.
func CalculateDurationTime(startTimer time.Time) string {
	return fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())
}

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
	case "doc":
		color = yellow
		icon = ""
	case "command":
		color = green
		icon = ""
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
	return fmt.Sprintf("%s%s%s", color, icon, color)
}

func SearchAndDeleteKeepFiles(rootPath string) error {

	var files []string

	err := filepath.Walk(rootPath, func(
		path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Base(path) == ".keep" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
