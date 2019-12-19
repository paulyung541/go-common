package web

import (
	"testing"
)

func Test_DoGet(t *testing.T) {
	m := make(map[string]interface{})
	if err := DoGet("https://api.github.com/users/torvalds", nil, &m); err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("%+v", m)
}
