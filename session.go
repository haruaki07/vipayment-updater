package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (v *VIPayment) AckSession() error {
	loginUrl := "https://vip-reseller.co.id/auth/login"
	res, err := v.Request(RequestOptions{
		"GET",
		loginUrl,
		nil,
		0, // no timeout
	})
	if err != nil {
		return fmt.Errorf("error get session: %v", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("error get session: %v", err)
	}

	if res.Request.URL.String() != loginUrl {
		// var base_url = 'https://vip-reseller.co.id/',
		//     csrf_key = 'thisisrandomhexstringaascoiqwjepq',
		//     ...
		js := doc.Find(`head script[type="text/javascript"]`).Last().Text()

		// @see https://regexr.com/729fj
		csrfRe := regexp.MustCompile("(?m)^\\s*(?:csrf_key)\\s*=\\s*(?:[\"'`])(.*)(?:[\"'`])")

		// []string{
		//   csrf_key = 'thisisrandomhexstringaascoiqwjepq',
		//   thisisrandomhexstringaascoiqwjepq
		// ]
		val := csrfRe.FindStringSubmatch(js)

		v.csrfToken = val[1]
	} else {
		val, exist := doc.Find("input#csrf_token").Attr("value")
		if !exist {
			return fmt.Errorf(
				"error get session: %v",
				errors.New("csrf token input element not found"),
			)
		}
		v.csrfToken = val

	}

	return nil
}
