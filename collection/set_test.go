package collection

import (
	"github.com/stretchr/testify/assert"
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

func TestDiff(t *testing.T) {
	s01 := NewSet()
	s02 := NewSet()
	assert.Equal(t, s01, Diff(s01, s02)) // no diff

	s03 := NewSet(1, 2, 3)
	s04 := NewSet()
	assert.Equal(t, s03, Diff(s03, s04)) // diff 1, 2, 3

	s05 := NewSet(1, 2, 3)
	s06 := NewSet(1, 2)
	assert.Equal(t, NewSet(3), Diff(s05, s06)) // diff 3

	s07 := NewSet(1, 2, 3)
	s08 := NewSet(1, 2, 3, 4)
	assert.Equal(t, NewSet(4), Diff(s07, s08)) // diff 4

	s09 := NewSet(1, 2, 3, 5)
	s10 := NewSet(1, 2, 3, 4)
	assert.Equal(t, NewSet(4, 5), Diff(s09, s10)) // diff 4, 5

	s11 := NewSet(1, 2, 3, 4, 5)
	s12 := NewSet(1, 2, 3, 6, 7)
	assert.Equal(t, NewSet(4, 5, 6, 7), Diff(s11, s12)) // diff 4, 5

	s13 := NewSet(1)
	s14 := NewSet(2)
	assert.Equal(t, NewSet(1, 2), Diff(s13, s14)) // no diff
}