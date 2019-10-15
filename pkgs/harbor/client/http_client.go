package harbor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func postLogin(url, user, passwd string) (cookie string) {
	httpClient := http.Client{}

	body := "principal=" + user + "&" + "password=" + passwd
	b := strings.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		fmt.Printf("http new request error: %v\n", err.Error())
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;param=value")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("2")
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("3")
		return ""
	}

	c := filterCookie(resp.Cookies())
	if c != "" {
		return c
	}
	return ""
}

func filterCookie(cookies []*http.Cookie) (c string) {
	for _, cookie := range cookies {
		parts := strings.Split(strings.TrimSpace(cookie.String()), ";")
		if len(parts) == 1 && parts[0] == "" {
			return ""
		}

		for _, part := range parts {
			cc := strings.Split(part, "=")
			if cc[0] == "sid" {
				return cc[1]
			}
		}
	}
	return ""
}

func get(url string) []byte {
	fmt.Printf("url:%s\n", url)
	httpClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic YWRtaW46SGFyYm9yMTIzNDU=")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http code: %d\n", resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	//fmt.Printf("body:%v\n", body)
	return body
}
