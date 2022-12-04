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
