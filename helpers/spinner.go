package helpers

import (
	"time"

	"github.com/briandowns/spinner"
)

//StartSpinner returns spinner reference
func StartSpinner(msg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = msg
	s.Color("green")
	s.Start()

	return s
}
