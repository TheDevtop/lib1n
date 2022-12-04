package lib1n

import (
	"testing"
)

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

func TestMap(t *testing.T) {
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

	nds := Map(ds, fn)

	check0 := nds["0"][0] == "T"
	check1 := nds["1"][1] == "T"
	check2 := nds["2"][2] == "T"

	if check0 && check1 && check2 {
		t.Log("Mapped successfully")
		return
	} else {
		t.Fail()
	}
}
