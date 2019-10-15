package harbor

import (
	"fmt"
	"testing"
)

func Test_postLogin(t *testing.T) {
	url := "http://exporter.harbor.com/c/login"
	fmt.Printf("%s\n", postLogin(url, "admin", "Harbor12345"))
}
