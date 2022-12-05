package lib1n

import (
	"testing"
)

func TestDecode(t *testing.T) {
	data := "foo=bar;baz\nos=plan9\ntext=lorem;ipsum"
	if ds, err := Decode([]byte(data)); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v\n", ds)
	}
}

func TestDecodeFail(t *testing.T) {
	data := "foo=bar;baz\nos=windows\ntext\nlorem=ipsum\n"
	if ds, err := Decode([]byte(data)); err == nil {
		t.Logf("ds: %v\n", ds)
		t.Fatal("This functions should have failed!\n")
	}
}

func TestEncode(t *testing.T) {
	data := DataSet{
		"foo":  []string{"bar", "baz"},
		"os":   []string{"plan9"},
		"text": []string{"lorem", "ipsum"},
	}
	str := string(Encode(data))
	t.Logf("%s\n", str)
}

func TestComplete(t *testing.T) {
	data := DataSet{
		"foo":  []string{"bar", "baz"},
		"os":   []string{"plan9"},
		"text": []string{"lorem", "ipsum"},
	}

	buf := Encode(data)
	t.Log("Encoded dataset\n")

	if data, err := Decode(buf); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v\n", data)
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
