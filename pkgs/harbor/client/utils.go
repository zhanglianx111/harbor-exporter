package harbor

import "os"

func getCookie() (cookie string) {
	return os.Getenv("cookie")
}
