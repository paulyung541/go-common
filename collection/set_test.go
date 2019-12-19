package collection

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet("abc", "def", "ght", "abc")
	if s.Len() != 3 {
		t.Fatal("s.Len() != 3")
	}
}

func TestSet2(t *testing.T) {
	s := NewSet("abc", "def", "ght")
	s.Add("aaa")

	if !s.All("aaa") {
		t.Fatal("dont have aaa")
	}
}

func TestSet3(t *testing.T) {
	s := NewSet("abc", "def", "ght")
	s.Add("aaa")

	if !s.Any("aaa", "bbb") {
		t.Fatal("dont have aaa")
	}
}