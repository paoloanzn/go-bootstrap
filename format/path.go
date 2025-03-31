package format

import (
	"fmt"
)

func FormatPath(path string) string {
	workingDirPrefix := "./"
	firstChar := path[0]

	switch firstChar {
	case '.':
		if path[1] == '/' {
			return path
		}

		formattedPath := fmt.Sprintf("%s%s", workingDirPrefix, path)
		return formattedPath
	case '/':
		return path
	default:
		formattedPath := fmt.Sprintf("%s%s", workingDirPrefix, path)
		return formattedPath
	}
}
