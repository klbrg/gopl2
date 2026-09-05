// Package greeting produces greetings.
package greeting

import (
	"fmt"

	"rsc.io/quote"
)

// Hello returns a greeting for name, followed by a proverb.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s! %s", name, quote.Go())
}
