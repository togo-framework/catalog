package catalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/togo-framework/togo"
)

func TestRegisterAndTables(t *testing.T) {
	var posted Table
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			_ = json.NewDecoder(r.Body).Decode(&posted)
			w.WriteHeader(200)
		case "GET":
			_, _ = w.Write([]byte(`{"tables":[{"catalog_name":"main","schema_name":"public","name":"orders"}]}`))
		}
	}))
	defer srv.Close()
	u := &unity{base: srv.URL, hc: srv.Client()}
	if err := u.Register(context.Background(), Table{Catalog: "main", Schema: "public", Name: "orders"}); err != nil {
		t.Fatal(err)
	}
	if posted.Name != "orders" {
		t.Fatalf("registered %+v", posted)
	}
	tbls, err := u.Tables(context.Background(), "main", "public")
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) != 1 || tbls[0].Name != "orders" {
		t.Fatalf("tables = %v", tbls)
	}
}

func TestFromKernelDefaultsToUnity(t *testing.T) {
	c := FromKernel(togo.New())
	if _, ok := c.(*unity); !ok {
		t.Fatalf("default catalog should be unity, got %T", c)
	}
}
