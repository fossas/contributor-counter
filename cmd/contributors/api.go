package main

import (
	"crypto/tls"
	"crypto/x509"
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

var client *http.Client

func newTLSConfig(insecure bool) *tls.Config {
	if !insecure {
		return &tls.Config{}
	}

	const dummyPEM = `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(dummyPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	return &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            roots,
	}
}

func NewClient(insecure bool) *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   newTLSConfig(insecure),
		},
	}
}

func Client(insecure bool) *http.Client {
	if client != nil {
		return client
	}
	c := NewClient(insecure)
	client = c
	return client
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
	res, err := Client(*flagInsecure).Do(req)
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
