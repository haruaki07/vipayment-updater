package main

import (
	"flag"
	"log"
)

var (
	username string
	password string
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("ERROR: %v\n", r)
		}
	}()

	parseFlags()

	vp := NewWithCredentials(username, password)

	err := vp.Login()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := vp.Logout()
		if err != nil {
			panic(err)
		}
	}()

	ip, err := GetMyPublicIP()
	if err != nil {
		panic(err)
	}

	ipAddress := []string{ip}
	err = vp.UpdateIPWhitelist(ipAddress)
	if err != nil {
		panic(err)
	}
}

func parseFlags() {
	flag.StringVar(&username, "user", "", "VIPayment account username/email/phone")
	flag.StringVar(&password, "pass", "", "VIPayment account password")

	flag.Parse()

	if username == "" {
		panic("Missing required parameter: -user")
	}

	if password == "" {
		panic("Missing required parameter: -pass")
	}
}
