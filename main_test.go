package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Server(t *testing.T) {
	req, err := http.NewRequest("GET", "http://some.host/weather/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	Server(res, req)

	act := res.Body.String()

	data := &ResponseJson{
		Weather: DataJson{},
	}
	err = json.Unmarshal([]byte(act), data)
	if err != nil {
		t.Fatalf("Expected correct JSON in response, responded with %s", act)
	}

}
