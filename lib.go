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

func Decode(buf []byte) (DataSet, error) {
	lns := strings.Split(string(buf), tokLn)
	ds := make(DataSet, len(lns))

	for i, ln := range lns {
		if ln == "" {
			continue
		}
		if eqLn := strings.Split(ln, tokEq); len(eqLn) < 2 {
			return nil, fmt.Errorf(errFmt, i)
		} else {
			df := []string(strings.Split(strings.Join(eqLn[1:], ""), tokSep))
			ds[eqLn[0]] = df
		}
	}
	return ds, nil
}

func Encode(ds DataSet) []byte {
	var str string
	for k, df := range ds {
		v := strings.Join(df, tokSep)
		str += fmt.Sprintf(dataFmt, k, v)
	}
	return []byte(str)
}
