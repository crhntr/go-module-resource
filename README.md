# Go Module Resource

A Concourse Resource to list Go module versions backed by the Go proxy.

## Source Configuration

```json
{
  "proxy": "http://localhost:8080",
  "module": "github.com/crhntr/neldermead"
}
```

The default value for "proxy" is https://proxy.golang.org.

The "direct" proxy value is not supported.

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