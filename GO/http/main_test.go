package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientUpperCase(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.UpperCase("anything")
	log.Println(err)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	// res: expected\r\n
	// due to the http protocol cleanup response
	res = strings.TrimSpace(res)
	log.Println(res)
}

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
