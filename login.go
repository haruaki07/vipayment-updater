package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (v *VIPayment) Login() (err error) {
	err = v.AckSession()
	if err != nil {
		return
	}

	data := &url.Values{}
	data.Set("user", v.Credentials.Username)
	data.Set("pass", v.Credentials.Password)
	data.Set("login", "true")
	data.Set("csrf_token", v.csrfToken)

	res, err := v.Request(RequestOptions{
		"POST",
		"https://vip-reseller.co.id/auth/login",
		strings.NewReader(data.Encode()),
		3 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("error login: %v", err)
	}
	defer res.Body.Close()

	return nil
}
