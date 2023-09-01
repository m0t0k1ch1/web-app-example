package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type APIClient struct {
	http.Client

	baseURL string
}

func NewAPIClient(baseURL string) (*APIClient, error) {
	return &APIClient{
		Client: http.Client{},

		baseURL: baseURL,
	}, nil
}

func (c *APIClient) DoAPI(ctx context.Context, method string, pathname string, params any, respBody any) (int, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse base url")
	}

	u.Path = path.Join(u.Path, pathname)

	var reqBody io.Reader

	switch method {
	case http.MethodGet:
		if params == nil {
			break
		}

		v, ok := params.(url.Values)
		if !ok {
			return 0, errors.New("failed to convert params into type url.Values")
		}

		u.RawQuery = v.Encode()

	case http.MethodPost:
		if params == nil {
			break
		}

		paramsBytes, err := json.Marshal(params)
		if err != nil {
			return 0, errors.Wrap(err, "failed to marshal params")
		}

		reqBody = bytes.NewBuffer(paramsBytes)

	default:
		return 0, errors.New("unsupported method")
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return 0, errors.Wrap(err, "failed to initialize request")
	}

	switch method {
	case http.MethodPost:
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if respBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			return 0, errors.Wrap(err, "failed to decode response body")
		}
	}

	return resp.StatusCode, nil
}
