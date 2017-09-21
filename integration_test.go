package functional

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestComposeCurried(t *testing.T) {
	var (
		sum     = func(a, b int) int { return a + b }
		mul     = func(a, b int) int { return a * b }
		dblSucc func(int) int
	)
	AssignComposed(&dblSucc,
		Apply(Curry(mul), 2),
		Apply(Curry(sum), 1))

	actual := dblSucc(9)
	if actual != 20 {
		t.Errorf("expected 20, got %d", actual)
	}
}

func TestDoReadAll(t *testing.T) {
	var (
		client = http.Client{}
		ts     = httptest.NewServer(
			http.HandlerFunc(
				func(rw http.ResponseWriter, req *http.Request) {
					fmt.Fprint(rw, "test server responding")
				}))

		readBody = func(resp *http.Response) ([]byte, error) {
			defer func() { _ = resp.Body.Close() }()
			return ioutil.ReadAll(resp.Body)
		}

		doReadAll func(string, string, io.Reader) ([]byte, error)
	)
	defer ts.Close()

	AssignComposed(&doReadAll,
		readBody,
		client.Do,
		http.NewRequest)

	actual, err := doReadAll(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Error(err)
	}

	if string(actual) != "test server responding" {
		t.Errorf("expected \"test server responding\", got \"%s\"", actual)
	}
}

func TestHttpGet(t *testing.T) {
	var (
		client = http.Client{}
		ts     = httptest.NewServer(
			http.HandlerFunc(
				func(rw http.ResponseWriter, req *http.Request) {
					fmt.Fprint(rw, "test server responding")
				}))

		readBody = func(resp *http.Response) ([]byte, error) {
			defer func() { _ = resp.Body.Close() }()
			return ioutil.ReadAll(resp.Body)
		}

		getReadAll func(string) ([]byte, error)
	)
	defer ts.Close()

	AssignComposed(&getReadAll,
		readBody,
		client.Do,
		Apply(
			Flip(
				Curry(
					Apply(
						Curry(http.NewRequest),
						http.MethodGet))),
			bytes.NewReader(nil)))

	actual, err := getReadAll(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if string(actual) != "test server responding" {
		t.Errorf("expected \"test server responding\", got \"%s\"", actual)
	}
}
