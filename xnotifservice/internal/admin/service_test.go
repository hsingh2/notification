package admin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func TestHTTP(t *testing.T) {

	svc := NewAdminService()
	//create endpoints
	endpoints := MakeEndpoints(svc)
	serverOptions := []kithttp.ServerOption{}
	//add handlers and spin server

	router := mux.NewRouter()
	AddAdminServiceRoutes(router, endpoints, serverOptions, log.NewJSONLogger(os.Stderr))
	server := httptest.NewServer(router)
	defer server.Close()

	for _, testcase := range []struct {
		method, url, body, want string
	}{
		{http.MethodGet, server.URL + "/health", ``, ``},
		{http.MethodGet, server.URL + "/token", ``, ``},
	} {
		request, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		response, _ := http.DefaultClient.Do(request)
		body, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(body))
	}
}
