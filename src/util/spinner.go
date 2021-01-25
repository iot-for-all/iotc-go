package util

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func NewSpinner(message string) *spinner.Spinner {
	spin := spinner.New(spinner.CharSets[9], 500*time.Millisecond, spinner.WithWriter(os.Stderr)) // Build our new spinner
	spin.Suffix = message
	spin.Start() // Start the spinner

	return spin
}