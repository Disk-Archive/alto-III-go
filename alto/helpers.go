package alto

import (
	"fmt"
	"strings"
)

// PipedStringToSlice parses the output returned from prometheus into a slice
func AltoStringToSlice(PipedString string) (parsedList []string, err error) {

	parsedList = strings.Split(PipedString, "|")

	if len(parsedList) < 1 {
		return nil, fmt.Errorf("prometheus returned empty output")
	}
	if parsedList[0] == "0" {
		return parsedList, fmt.Errorf("error processing command")
	}

	return
}
