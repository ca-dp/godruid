package godruid

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DefaultRetentionRulesEndPoint   = "/druid/coordinator/v1/rules/%s"
	DefaultCompactionConfigEndPoint = "/druid/coordinator/v1/config/compaction"
)

var (
	ErrDatasourceEmpty = errors.New("coordinatorClient: datasource is empty")
)

type CoordinatorClient struct {
	Url      string
	EndPoint string

	Debug        bool
	LastRequest  string
	LastResponse string
	HttpClient   *http.Client
}

func (c *CoordinatorClient) CreateOrUpdateRetentionRule(ctx context.Context, datasource string, rules RetentionRulesSpec, authToken string) (err error) {
	if c.EndPoint == "" {
		if datasource == "" {
			return ErrDatasourceEmpty
		}
		c.EndPoint = fmt.Sprintf(DefaultRetentionRulesEndPoint, datasource)
	}
	var reqJson []byte
	if c.Debug {
		reqJson, err = json.MarshalIndent(rules.(*RetentionRules).Rules, "", "  ")
	} else {
		reqJson, err = json.Marshal(rules.(*RetentionRules).Rules)
	}
	if err != nil {
		return
	}

	_, err = c.Post(ctx, reqJson, authToken)
	return err
}

func (c *CoordinatorClient) CreateOrUpdateCompactionConfig(ctx context.Context, compactionConfig *CompactionConfig, authToken string) (err error) {
	if c.EndPoint == "" {
		c.EndPoint = DefaultCompactionConfigEndPoint
	}
	var reqJson []byte
	if c.Debug {
		reqJson, err = json.MarshalIndent(compactionConfig, "", "  ")
	} else {
		reqJson, err = json.Marshal(compactionConfig)
	}
	if err != nil {
		return
	}

	_, err = c.Post(ctx, reqJson, authToken)
	return err
}

func (c *CoordinatorClient) Post(ctx context.Context, req []byte, authToken string) (result []byte, err error) {
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
