package main

import (
	"log"
	"net/http/cookiejar"
	"os"
)

type VIPaymentCredentials struct {
	Username string
	Password string
}

type VIPayment struct {
	Credentials VIPaymentCredentials
	Jar         *cookiejar.Jar
	logger      *log.Logger
	csrfToken   string
}

func NewWithCredentials(user, pass string) *VIPayment {
	jar, _ := cookiejar.New(nil)
	logger := log.New(os.Stdout, "", log.Ltime|log.Lshortfile)

	return &VIPayment{
		VIPaymentCredentials{
			user,
			pass,
		},
		jar,
		logger,
		"",
	}
}
