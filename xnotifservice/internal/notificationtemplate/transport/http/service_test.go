package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/cockroachdb"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/implementation"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func TestHTTP(t *testing.T) {
	repo := cockroachdb.MockRepository()
	svc := implementation.NewTemplateService(repo, log.NewJSONLogger(os.Stderr))
	//create endpoints
	endpoints := transport.MakeEndpoints(svc)
	serverOptions := []kithttp.ServerOption{}
	//add handlers and spin server

	router := mux.NewRouter()
	NewNotificationTemplateService(router, endpoints, serverOptions, log.NewJSONLogger(os.Stderr))
	server := httptest.NewServer(router)
	defer server.Close()

	for _, testcase := range []struct {
		method, url, body, want string
	}{
		{http.MethodPost, server.URL + "/api/v1/notificationTemplates", `{"Name":"Harvinder"}`, `{"v":"12"}`},
		{http.MethodGet, server.URL + "/api/v1/notificationTemplates/count/", ``, `100`},
		{http.MethodGet, server.URL + "/api/v1/notificationTemplates/dadsd", ``, `100`},
	} {
		request, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		response, _ := http.DefaultClient.Do(request)
		body, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(body))
	}
}
