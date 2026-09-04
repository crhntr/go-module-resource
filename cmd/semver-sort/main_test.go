package main

import (
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	for _, tt := range []struct {
		name    string
		in      []string
		skipPre bool
		want    []string
	}{
		{
			name: "sorts ascending",
			in:   []string{"v1.2.0", "v1.10.0", "v1.9.0"},
			want: []string{"v1.2.0", "v1.9.0", "v1.10.0"},
		},
		{
			name:    "drops prereleases when asked",
			in:      []string{"v1.0.0", "v1.1.0-rc.1"},
			skipPre: true,
			want:    []string{"v1.0.0"},
		},
		{
			name: "keeps prereleases otherwise",
			in:   []string{"v1.0.0", "v1.1.0-rc.1"},
			want: []string{"v1.0.0", "v1.1.0-rc.1"},
		},
		{
			name: "discards invalid lines",
			in:   []string{"v1.0.0", "not-a-version", ""},
			want: []string{"v1.0.0"},
		},
		{
			// github.com/mattn/go-sqlite3 has both. `go list -m -versions` and
			// `go get @latest` both report v1.14.50; the +incompatible tags are
			// major versions published without module support and are never
			// what a dependency bump should move to.
			name:    "drops +incompatible even though it sorts highest",
			in:      []string{"v1.14.50", "v2.0.0+incompatible", "v2.0.3+incompatible"},
			skipPre: true,
			want:    []string{"v1.14.50"},
		},
		{
			// +incompatible is build metadata, not a prerelease, so the
			// prerelease filter never sees it. It has to be dropped whether or
			// not prereleases are being kept.
			name: "drops +incompatible while keeping prereleases",
			in:   []string{"v1.14.50", "v1.15.0-rc.1", "v2.0.3+incompatible"},
			want: []string{"v1.14.50", "v1.15.0-rc.1"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var out strings.Builder
			if err := sort(&out, strings.NewReader(strings.Join(tt.in, "\n")), tt.skipPre); err != nil {
				t.Fatalf("sort() error = %v", err)
			}
			got := strings.Fields(out.String())
			if strings.Join(got, ",") != strings.Join(tt.want, ",") {
				t.Errorf("sort(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
