package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type RequestOptions struct {
	Method  string
	Url     string
	Body    io.Reader
	Timeout time.Duration
}

func (v *VIPayment) Request(opts RequestOptions) (*http.Response, error) {
	client := &http.Client{
		Jar:     v.Jar,
		Timeout: opts.Timeout,
	}

	req, err := http.NewRequest(opts.Method, opts.Url, opts.Body)
	if err != nil {
		return nil, fmt.Errorf("error init request: %v", err)
	}

	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	v.Jar.SetCookies(res.Request.URL, res.Cookies())

	return res, nil
}
