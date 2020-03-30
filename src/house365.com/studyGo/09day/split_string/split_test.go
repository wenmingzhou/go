package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcbc", "b")
	want := []string{"a", "c", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got :%v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got :%v\n", want, got)
	}
}

func Test3Split(t *testing.T) {
	got := Split("abcdef", "bc")
	want := []string{"a", "def"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got :%v\n", want, got)
	}
}
