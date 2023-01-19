package discord

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func PostToDiscordWebHook(webhook_url string, params map[string]string) string {
	cli := &http.Client{Timeout: time.Duration(30) * time.Second}
	values := url.Values{}

	for key, element := range params {
		values.Add(key, element)
	}

	rsp, err := cli.PostForm(webhook_url, values)
	if err != nil {
		return err.Error()
	} else if rsp.StatusCode != 204 {
		return rsp.Status
	}

	fmt.Println(rsp.Status)
	defer rsp.Body.Close()
	return ""
}
