package lib1n

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	data := "0=X;Y;Z\n1=X;Y;Z\n2=X;Y;Z\n"
	var (
		ds  DataSet
		err error
	)

	if ds, err = Decode([]byte(data)); err != nil {
		t.Fatal(err)
	}

	check0 := ds["0"][0] == "X"
	check1 := ds["1"][1] == "Y"
	check2 := ds["2"][2] == "Z"

	if check0 && check1 && check2 {
		t.Log("Decoded successfully")
		return
	} else {
		t.Fail()
	}
}

func TestDecodeFail(t *testing.T) {
	data := "0=X;Y;Z\nbroken\n2=X;Y;Z\n"

	var (
		ds  DataSet
		err error
	)

	if ds, err = Decode([]byte(data)); err == nil {
		t.Fail()
	}

	if ds == nil {
		t.Log("Failed successfully")
		return
	} else {
		t.Fail()
	}
}

func TestEncode(t *testing.T) {
	data := "0=X;O;O\n1=O;X;O\n2=O;O;X\n"
	ds := DataSet{
		"0": {"X", "O", "O"},
		"1": {"O", "X", "O"},
		"2": {"O", "O", "X"},
	}

	if string(Encode(ds)) != data {
		t.Fail()
	}
}

func TestGoEncode(t *testing.T) {
	ds := DataSet{
		"0": {"X", "O", "O"},
		"1": {"O", "X", "O"},
		"2": {"O", "O", "X"},
	}

	data := string(GoEncode(ds))

	check0 := strings.Contains(data, "0=X;O;O\n")
	check1 := strings.Contains(data, "1=O;X;O\n")
	check2 := strings.Contains(data, "2=O;O;X\n")

	if check0 && check1 && check2 {
		t.Log("Encoded successfully")
		return
	} else {
		t.Fail()
	}
}

func TestGoDecode(t *testing.T) {
	data := "0=X;Y;Z\n1=X;Y;Z\n2=X;Y;Z\n"
	ds := GoDecode([]byte(data))

	check0 := ds["0"][0] == "X"
	check1 := ds["1"][1] == "Y"
	check2 := ds["2"][2] == "Z"

	if check0 && check1 && check2 {
		t.Log("Decoded successfully")
		return
	} else {
		t.Fail()
	}
}

func TestGoDecodeFail(t *testing.T) {
	data := "0=X;Y;Z\nbroken\n2=X;Y;Z\n"
	ds := GoDecode([]byte(data))

	check0 := ds["0"][0] == "X"
	_, check1 := ds["1"]
	check2 := ds["2"][2] == "Z"

	if check0 && !check1 && check2 {
		t.Log("Failed successfully")
		return
	} else {
		t.Fail()
	}
}
