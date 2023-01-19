package discord

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func PostToDiscordWebHook(wenhook_url string, params map[string]string) {
	cli := &http.Client{Timeout: time.Duration(30) * time.Second}
	values := url.Values{}

	for key, element := range params {
		values.Add(key, element)
	}

	rsp, err := cli.PostForm(wenhook_url, values)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rsp.Body.Close()

	body, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(body))
}
