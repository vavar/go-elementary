package apiproxy_test

import (
	"github.com/vavar/go-elementary/apiproxy"
	"testing"
	"net/http"
)

func TestList(t *testing.T) {
	apiproxy.List(&http.Client{},"https://reqres.in/api/users?page=2")
}