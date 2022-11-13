package main

import (
	"flag"
	"log"
)

var (
	username string
	password string
)

func init() {
	flag.StringVar(&username, "user", "", "VIPayment account username/email/phone")
	flag.StringVar(&password, "pass", "", "VIPayment account password")

	flag.Parse()

	if username == "" {
		log.Panicln("Missing required parameter: -user")
	}

	if password == "" {
		log.Panicln("Missing required parameter: -pass")
	}
}

func main() {
	vp := NewWithCredentials(username, password)

	err := vp.Login()
	if err != nil {
		log.Fatalln(err)
	}

	ip, err := GetMyPublicIP()
	if err != nil {
		log.Fatalln(err)
	}

	ipAddress := []string{ip}
	err = vp.UpdateIPWhitelist(ipAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
