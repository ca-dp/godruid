package godruid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DefaultSupervisorEndPoint = "/druid/indexer/v1/supervisor"
)

type OverlordClient struct {
	Url      string
	EndPoint string

	Debug        bool
	LastRequest  string
	LastResponse string
	HttpClient   *http.Client
}

func (c *OverlordClient) CreateOrUpdateSupervisor(ctx context.Context, spec SupervisorSpec, authToken string) (err error) {
	if c.EndPoint == "" {
		c.EndPoint = DefaultSupervisorEndPoint
	}
	var reqJson []byte
	if c.Debug {
		reqJson, err = json.MarshalIndent(spec, "", "  ")
	} else {
		reqJson, err = json.Marshal(spec)
	}
	if err != nil {
		return
	}

	result, err := c.Post(ctx, reqJson, authToken)
	if err != nil {
		return
	}

	return spec.onResponse(result)
}

func (c *OverlordClient) Post(ctx context.Context, req []byte, authToken string) (result []byte, err error) {
	endPoint := c.EndPoint
	if c.Debug {
		endPoint += "?pretty"
		c.LastRequest = string(req)
	}
	if err != nil {
		return
	}

	request, err := http.NewRequest("POST", c.Url+endPoint, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		cookie := &http.Cookie{
			Name:  "skylight-aaa",
			Value: authToken,
		}
		request.AddCookie(cookie)
	}
	resp, err := c.HttpClient.Do(request.WithContext(ctx))
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if c.Debug {
		c.LastResponse = string(result)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(result))
	}

	return
}
