package common

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type RequestHead struct {
	SessionID string
}

type Request[T any] struct {
	Head RequestHead
	Body T
}

func DecodeRequest[T any](r *http.Request) (Request[T], error) {
	req := Request[T]{
		Head: RequestHead{
			SessionID: r.Header.Get("X-Session-Id"),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil && !errors.Is(err, io.EOF) {
		return req, err
	}
	return req, nil
}
