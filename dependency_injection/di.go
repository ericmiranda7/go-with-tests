package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, msg string) {
	_, err := fmt.Fprintf(w, msg)
	if err != nil {
		return
	}
}
