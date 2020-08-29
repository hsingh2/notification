package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	template "cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate"
	"cto-github.cisco.com/NFV-BU/xnotifservice/internal/notificationtemplate/transport"
)

//Error
var (
	ErrBadRouting = errors.New("bad routing")
)

// NewNotificationTemplateService wires Go kit endpoints to the HTTP transport.
func NewNotificationTemplateService(r *mux.Router, svcEndpoints transport.Endpoints, options []kithttp.ServerOption) {
	// set-up router and initialize http endpoints

	// HTTP Post - /api/v1/notificationTemplates
	r.Methods(http.MethodPost).Path("/api/v1/notificationTemplates").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeCreateResponse,
		options...,
	))

	// HTTP Post - /api/v1/notificationTemplates/{id}
	r.Methods(http.MethodGet).Path("/api/v1/notificationTemplates/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodeGetByIDRequest,
		encodeGetByIDResponse,
		options...,
	))

	// HTTP Get - /api/v1/notificationTemplates/count
	r.Methods(http.MethodGet).PathPrefix("/api/v1/notificationTemplates/count/").Handler(kithttp.NewServer(
		svcEndpoints.Count,
		func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
		encodeCountResponse,
		options...,
	))
	return
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateNotificationTemplateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.NotificationTemplate); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.GetNotificationTemplateByIDRequest{ID: id}, nil
}

func encodeCountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeGetByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		//these are business logic errors
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
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
	case template.ErrNotificationNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
