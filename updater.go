package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (v *VIPayment) UpdateIPWhitelist(ipAddress []string) (err error) {
	err = v.AckSession()
	if err != nil {
		return
	}

	data := &url.Values{}
	data.Set("csrf_token", v.csrfToken)
	data.Set("ipstat", strings.Join(ipAddress, ","))
	data.Set("apistatus", "production")
	data.Set("changeapi", "true")

	res, err := v.Request(RequestOptions{
		"POST",
		"https://vip-reseller.co.id/account/profile",
		strings.NewReader(data.Encode()),
		3 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("error update ip whitelist: %v", err)
	}
	defer res.Body.Close()

	return nil
}
