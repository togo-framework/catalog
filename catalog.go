// Package catalog is togo's open-catalog capability — register and discover
// tables in an open governance catalog (Unity Catalog / Iceberg REST), so a
// togo app's data becomes governed + discoverable by any engine (Spark, Trino,
// Databricks) with zero lock-in. The default `unity` backend speaks the Unity
// Catalog OSS REST API. Select/point it with `togo provider:use catalog unity`.
//
//   CATALOG_URL    catalog server base URL (default http://localhost:8080)
//   CATALOG_TOKEN  bearer token (secret)
package catalog

import (
	"context"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// Table is a catalog table entry.
type Table struct {
	Catalog  string `json:"catalog_name"`
	Schema   string `json:"schema_name"`
	Name     string `json:"name"`
	Location string `json:"storage_location,omitempty"`
	Format   string `json:"data_source_format,omitempty"`
}

// Catalog registers and lists tables in an open catalog.
type Catalog interface {
	Register(ctx context.Context, t Table) error
	Tables(ctx context.Context, catalog, schema string) ([]Table, error)
}

func init() {
	togo.RegisterProviderFunc("catalog", togo.PriorityService, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCatalog, "unity", newUnity(k), true)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "catalog")
		}
		return nil
	})
}

// FromKernel returns the active Catalog backend (or nil).
func FromKernel(k *togo.Kernel) Catalog {
	if v, ok := k.Get(providers.CapCatalog); ok {
		if c, ok := v.(Catalog); ok {
			return c
		}
	}
	return nil
}
