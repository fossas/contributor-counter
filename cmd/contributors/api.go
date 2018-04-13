package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type PagedResponse struct {
	IsLastPage    bool
	NextPageStart int
	Values        []interface{}
}

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

func GetPaged(endpoint *url.URL, user, password string) ([]interface{}, error) {
	var values []interface{}
	page, err := GetSinglePage(endpoint.String(), user, password)
	if err != nil {
		return nil, err
	}
	values = append(values, page.Values...)
	for !page.IsLastPage {
		query := url.Values{}
		query.Set("start", strconv.Itoa(page.NextPageStart))
		endpoint.RawQuery = query.Encode()
		page, err = GetSinglePage(endpoint.String(), user, password)
		if err != nil {
			return nil, err
		}
		values = append(values, page.Values...)
	}
	return values, nil
}

func GetSinglePage(endpoint, user, password string) (PagedResponse, error) {
	// Send request.
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return PagedResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, password)
	debugf("API request: %#v", endpoint)
	res, err := client.Do(req)
	if err != nil {
		return PagedResponse{}, err
	}
	defer res.Body.Close()

	// Parse response.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return PagedResponse{}, err
	}
	debugf("API response: %#v", string(body))
	var page PagedResponse
	err = json.Unmarshal(body, &page)
	if err != nil {
		return PagedResponse{}, err
	}

	return page, nil
}
