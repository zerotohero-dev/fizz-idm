package downstream

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"io/ioutil"
	"net/http"
)

func decodeRelayWelcomeEmailResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.RelayWelcomeEmailResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeRelayEmailVerificationEmailResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.RelayEmailVerifiedEmailResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeRelayEmailVerifiedEmailResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.RelayEmailVerifiedEmailResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeTokenCreateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.TokenCreateResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeHashCreateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.HashCreateResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeJwtCreateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.JwtCreateResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeJwtVerifyResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.JwtVerifyResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeHashVerifyResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response reqres.HashVerifyResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
