<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/catalog</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/catalog"><img src="https://pkg.go.dev/badge/github.com/togo-framework/catalog.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/catalog
```

<!-- /togo-header -->

togo's **catalog** capability — register + discover tables in an open governance
catalog (**Unity Catalog** / Iceberg REST), so your app data is governed and
readable by any engine (Spark, Trino, Databricks) with zero lock-in.

```bash
togo install togo-framework/catalog
togo config:set CATALOG_URL http://localhost:8080
```

```go
c := catalog.FromKernel(k)
_ = c.Register(ctx, catalog.Table{Catalog: "main", Schema: "public", Name: "orders"})
```

| Key | Purpose |
|---|---|
| `CATALOG_URL` | Unity Catalog server |
| `CATALOG_TOKEN` | bearer token (secret) |
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
