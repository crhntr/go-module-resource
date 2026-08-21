# Go Module Resource

A custom [Concourse](https://concourse-ci.org) [Resource Type](https://concourse-ci.org/docs/resource-types/)
to list [Go module](https://go.dev/ref/mod#go-mod-file-ident) versions backed by the
[Go proxy](https://go.dev/ref/mod#module-proxy).

## Source Configuration

Source configuration fields:
- `proxy`: The Go module proxy (default https://proxy.golang.org).
  The "direct" proxy value is not supported.
- `module`: The Go module name that you want to fetch.

### Example

See ./ci/example.yml

```yaml
resource_types:
  - name: go-module
    type: registry-image
    source:
      repository: ghcr.io/crhntr/go-module-resource
      tag: latest
      username: your-username
      password: ((your-github-pat))

resources:
  - name: dependency-mod
    type: go-module
    check_every: 24h
    source:
      module: golang.org/x/mod

jobs: []
```

## Behavior

`check`: List module versions in the Go proxy.

The newline separated versions are correctly sorted.

`get`: Download Module Version Info

Writes 4 files:
- go.mod: The module version's mod file.
- info.json: Metadata about the version.
- version.txt: A plain text file with the version.
- go-get.sh: A script that calls `go get example.com@v0.1.2`
  where the module path and version come from the config and version respectively.

`put`: **Not implemented**

This does not make sense for this resource type.
