package lib1n

import (
	"fmt"
	"strings"
)

const (
	tokLn   = "\n"
	tokEq   = "="
	tokSep  = ";"
	dataFmt = "%s=%s\n"
	errFmt  = "ambiguous tokens at index %d"
)

type DataSet map[string][]string

// Decode data to DataSet
// Warning: Concurrent function
func GoDecode(buf []byte) DataSet {
	type result struct {
		key  string
		vals []string
	}

	lns := strings.Split(string(buf), tokLn)
	chResult := make(chan result, len(lns))
	nds := make(DataSet, len(lns))

	// Try to decode, return empty is fails
	for _, ln := range lns {
		go func(l string) {
			if l == "" {
				chResult <- result{key: "", vals: nil}
				return
			}
			if eqln := strings.Split(l, tokEq); len(eqln) < 2 {
				chResult <- result{key: "", vals: nil}
				return
			} else {
				k := eqln[0]
				v := []string(strings.Split(strings.Join(eqln[1:], ""), tokSep))

				// Sends result trough channel
				chResult <- result{key: k, vals: v}
				return
			}
		}(ln)

		// Collect results, and return new dataset
		res := <-chResult
		if res.key != "" && res.vals != nil {
			nds[res.key] = res.vals
		}
	}

	return nds
}

// Decode data to DataSet
func Decode(buf []byte) (DataSet, error) {
	lns := strings.Split(string(buf), tokLn)
	ds := make(DataSet, len(lns))

	for i, ln := range lns {
		if ln == "" {
			continue
		}
		if eqln := strings.Split(ln, tokEq); len(eqln) < 2 {
			return nil, fmt.Errorf(errFmt, i)
		} else {
			df := []string(strings.Split(strings.Join(eqln[1:], ""), tokSep))
			ds[eqln[0]] = df
		}
	}
	return ds, nil
}

// Encode DataSet to slice
func Encode(ds DataSet) []byte {
	var str string
	for k, df := range ds {
		v := strings.Join(df, tokSep)
		str += fmt.Sprintf(dataFmt, k, v)
	}
	return []byte(str)
}
