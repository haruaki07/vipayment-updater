package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetMyPublicIP() (ip string, err error) {
	res, err := http.Get("https://icanhazip.com")
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("get ip response status non-ok: %d", res.StatusCode)
		return
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	ip = string(bodyBytes)
	return
}
