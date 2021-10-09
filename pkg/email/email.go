package email

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

// GetSendHandler ...
func GetSendHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		ep,
		decodeSendRequest,
		encodeSendResponse,
		options...,
	)
}

// decodeSendRequest
func decodeSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req sendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeSendResponse
func encodeSendResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(sendResponse)
	if !ok {
		return errors.New("error decoding")
	}
	return json.NewEncoder(w).Encode(resp)
}
