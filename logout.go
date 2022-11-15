package main

import "time"

func (v *VIPayment) Logout() (err error) {
	err = v.AckSession()
	if err != nil {
		return
	}

	res, err := v.Request(RequestOptions{
		"GET",
		"https://vip-reseller.co.id/auth/logout",
		nil,
		3 * time.Second,
	})
	if err != nil {
		return
	}
	defer res.Body.Close()

	return nil
}
