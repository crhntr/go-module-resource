package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"slices"
	"strings"

	"golang.org/x/mod/semver"
)

func main() {
	skipPre := true
	flag.BoolVar(&skipPre, "no-prerelease", false, "drop prerelease versions")
	flag.Parse()
	if err := sort(os.Stdout, os.Stdin, skipPre); err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func sort(w io.Writer, r io.Reader, skipPre bool) error {
	var vs []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		v := strings.TrimSpace(sc.Text())
		if !semver.IsValid(v) {
			continue
		}
		if skipPre && semver.Prerelease(v) != "" {
			continue
		}
		// A +incompatible version is a v2 or higher tag on a module published
		// without module support. The go command does not offer them once the
		// module has a go.mod, and moving to one is a major version change
		// rather than an update, so they are never a version to bump to.
		//
		// This has to be its own check: +incompatible is build metadata, not a
		// prerelease, so the filter above never sees it while semver.Sort
		// happily ranks it above every v0 and v1 release.
		if semver.Build(v) != "" {
			continue
		}
		vs = append(vs, v)
	}
	if err := sc.Err(); err != nil {
		return err
	}
	semver.Sort(vs)
	vs = slices.Compact(vs)
	for _, v := range vs {
		if _, err := io.WriteString(w, v+"\n"); err != nil {
			return err
		}
	}
	return nil
}
