package http

import "fmt"

func FormatSslString(useSsl bool, uri string) string {
	if useSsl {
		return fmt.Sprintf("https://%s", uri)
	}
	return fmt.Sprintf("http://%s", uri)
}
