package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSayHello(t *testing.T) {
	// startup the server and make sure its response starts with "Hello"
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	sayHello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("sayHello should return 200, got %d", resp.StatusCode)
	}
	if !strings.HasPrefix(string(body), "Hello") {
		t.Error("sayHello response body should start with Hello")
	}
	if !strings.Contains(string(body), "Env Vars") {
		t.Error("sayHello response body should contain an 'Env Vars' section")
	}
}
