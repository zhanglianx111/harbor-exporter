package harbor

import (
	"crypto/tls"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var request = gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func postLogin(url, user, passwd string) (cookie string) {
	httpClient := http.Client{}

	body := "principal=" + user + "&" + "password=" + passwd
	b := strings.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		log.Errorf("http new request error: %v\n", err.Error())
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;param=value")
	resp, err := httpClient.Do(req)
	if err != nil {
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("response code: %d\n", resp.StatusCode)
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
	log.Infof("url:%s\n", url)

	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("request method:%s, url: %s\n", http.MethodGet, url)
		return nil
	}
	req.Header.Add("accept", "application/json")
	baseAuth := "Basic " + os.Getenv("baseAuth")
	req.Header.Add("authorization", baseAuth)

	/* TODO 使用cookie */
	//req.Header.Add("cookie", os.Getenv("cookie"))

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("get respones for harbor err: %v\n", err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Errorf("respone's code: %d\n", resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read respone body err: %v\n", err.Error())
		return nil
	}

	return body
}
