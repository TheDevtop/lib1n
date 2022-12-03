package lib1n

import (
	"fmt"
	"strings"
)

const (
	tokLn   = "\n"
	tokKey  = "="
	tokSep  = ";"
	dataFmt = "%s=%s\n"
	errFmt  = "ambiguous data at line %d"
)

type DataFrame []string
type DataSet map[string]DataFrame

func Decode(buf []byte) (DataSet, error) {
	lns := strings.Split(string(buf), tokLn)
	ds := make(DataSet, len(lns))

	for i, ln := range lns {
		rln := strings.Split(ln, tokKey)
		if len(rln) < 2 {
			return nil, fmt.Errorf(errFmt, i)
		}
		df := DataFrame(strings.Split(strings.Join(rln[1:], ""), tokSep))
		ds[rln[0]] = df
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
