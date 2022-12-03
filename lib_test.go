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
	data := "foo=bar;baz\nos=windows\ntext=lorem;ipsum=lmao\n"
	if _, err := Decode([]byte(data)); err == nil {
		t.Fatal("This functions should have failed!\n")
	} else {
		t.Log(err)
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
