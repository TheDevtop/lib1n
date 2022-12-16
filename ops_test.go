package lib1n

import (
	"strconv"
	"testing"
)

// Test find function
// Find the letter 'X'
func TestFind(t *testing.T) {
	ds := DataSet{
		"0": {"X", "O", "O"},
		"1": {"O", "O", "O"},
		"2": {"O", "O", "X"},
	}

	nds := Find(ds, "X")

	_, check0 := nds["0"]
	_, check1 := nds["1"]
	_, check2 := nds["2"]

	if check0 && !check1 && check2 {
		t.Log("Pattern found at the right indeces")
		return
	} else {
		t.Fail()
	}
}

// Test replace function
// Replace all 'X' with 'T'
func TestReplace(t *testing.T) {
	ds := DataSet{
		"0": {"X", "O", "O"},
		"1": {"O", "X", "O"},
		"2": {"O", "O", "X"},
	}

	nds := Replace(ds, "X", "T")

	check0 := nds["0"][0] == "T"
	check1 := nds["1"][1] == "T"
	check2 := nds["2"][2] == "T"

	if check0 && check1 && check2 {
		t.Log("Pattern found, and replaced at the right indeces")
		return
	} else {
		t.Fail()
	}
}

// Test map function
// Perform higher-order replace on data, replace all 'X' with 'T'
func TestMapFilter(t *testing.T) {
	ds := DataSet{
		"0": {"X", "O", "O"},
		"1": {"O", "X", "O"},
		"2": {"O", "O", "X"},
	}

	var fn = func(df []string) []string {
		for fi, fv := range df {
			if fv == "X" {
				df[fi] = "T"
			}
		}
		return df
	}

	nds := MapFilter(ds, fn)

	check0 := nds["0"][0] == "T"
	check1 := nds["1"][1] == "T"
	check2 := nds["2"][2] == "T"

	if check0 && check1 && check2 {
		t.Log("Map/Filter applied successfully")
		return
	} else {
		t.Fail()
	}
}

// Test clean functions
// Remove ambiguous tokens
func TestClean(t *testing.T) {
	ds := DataSet{
		"0":  {"X", "O", "O"},
		"1":  {"O", "X;", "O"},
		"2=": {"O", "O", "X"},
	}

	nds := Clean(ds)

	check0 := nds["0"][0] == "X"
	check1 := nds["1"][1] == "X"
	check2 := nds["2"][2] == "X"

	if check0 && check1 && check2 {
		t.Log("Cleaned succesfully")
		return
	} else {
		t.Fail()
	}
}

// Test reduce function
// Perform higher-order sum on data
func TestReduce(t *testing.T) {
	ds := DataSet{
		"0": {"1", "2", "3"}, // Sum = 6
		"1": {"4", "5", "6"}, // Sum = 15
		"2": {"7", "8", "9"}, // Sum = 24
	}

	var fn = func(key string, df []string) string {
		var r int
		for _, v := range df {
			if ir, er := strconv.Atoi(v); er != nil {
				ir = 0
			} else {
				r += ir
			}
		}
		return strconv.Itoa(r)
	}

	df := Reduce(ds, fn)

	check0 := df[0] == "6"
	check1 := df[1] == "15"
	check2 := df[2] == "24"

	if check0 && check1 && check2 {
		t.Log("Reduce applied successfully")
		return
	} else {
		t.Fail()
	}
}
