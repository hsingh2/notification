package admin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//Error
var (
	ErrBadRouting = errors.New("bad routing")
)

// AddAdminServiceRoutes ...
func AddAdminServiceRoutes(router *mux.Router, scvEndpoint Endpoints, options []kithttp.ServerOption) {
	// HTTP Get - /health
	router.Methods(http.MethodGet).Path("/notification/admin/health").Handler(kithttp.NewServer(
		scvEndpoint.Health,
		decodeHealthRequest,
		encodeHealthResponse,
		options...,
	))

	// HTTP GET - /token
	router.Methods(http.MethodGet).Path("/token").Handler(kithttp.NewServer(
		scvEndpoint.Token,
		decodeTokenRequest,
		encodeTokenResponse,
		options...,
	))

	return
}

func encodeHealthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHealthRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return nil, nil
}

func encodeTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeTokenRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return nil, nil
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
