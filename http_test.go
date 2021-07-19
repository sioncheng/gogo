package main

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

type RemoteSolver struct {
	MathServerUrl string
	Client        *http.Client
}

func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		rs.MathServerUrl+"?expression="+url.QueryEscape((expression)),
		nil)
	if err != nil {
		return 0, err
	}

	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(contents))
	}

	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

type info struct {
	expression string
	code       int
	body       string
}

var io info

func TestHttp(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expression")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression: " + expression))
			} else {
				rw.WriteHeader(io.code)
				rw.Write([]byte(io.body))
			}
		}))

	defer server.Close()

	rs := RemoteSolver{
		MathServerUrl: server.URL,
		Client:        server.Client(),
	}

	data := []struct {
		name   string
		io     info
		result float64
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22},
		{"case2", info{"2 + 2", http.StatusOK, "4"}, 4},
	}

	for _, d := range data {
		t.Run(d.name, func(tt *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				tt.Errorf("io %f, go %f", d.result, result)
			}
			if err != nil {
				tt.Error(err.Error())
			}
		})
	}
}
