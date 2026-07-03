package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// unity speaks the Unity Catalog OSS REST API.
type unity struct {
	base  string
	token string
	hc    *http.Client
}

func newUnity(k *togo.Kernel) *unity {
	return &unity{
		base:  strings.TrimRight(providers.Value(k, providers.CapCatalog, "unity", "url", "http://localhost:8080", false), "/"),
		token: providers.Value(k, providers.CapCatalog, "unity", "token", "", true),
		hc:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (u *unity) req(ctx context.Context, method, path string, body any) (*http.Response, error) {
	var r *bytes.Reader = bytes.NewReader(nil)
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req, _ := http.NewRequestWithContext(ctx, method, u.base+path, r)
	if u.token != "" {
		req.Header.Set("Authorization", "Bearer "+u.token)
	}
	req.Header.Set("Content-Type", "application/json")
	return u.hc.Do(req)
}

func (u *unity) Register(ctx context.Context, t Table) error {
	resp, err := u.req(ctx, "POST", "/api/2.1/unity-catalog/tables", t)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("catalog register: %s", resp.Status)
	}
	return nil
}

func (u *unity) Tables(ctx context.Context, catalog, schema string) ([]Table, error) {
	q := url.Values{"catalog_name": {catalog}, "schema_name": {schema}}
	resp, err := u.req(ctx, "GET", "/api/2.1/unity-catalog/tables?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("catalog list: %s", resp.Status)
	}
	var out struct {
		Tables []Table `json:"tables"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out.Tables, nil
}
