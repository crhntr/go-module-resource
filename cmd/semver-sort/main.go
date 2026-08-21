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
