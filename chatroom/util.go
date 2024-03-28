package chatroom

import "strings"

func trimInput(input *string) string {
	return strings.Trim(*input, "\n")
}
