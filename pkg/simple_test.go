package pkg_test

import (
  "testing"
  "github.com/tonygilkerson/gosimplepac/pkg"
)

func TestGetName(t *testing.T) {
	want := "I am simplex"
	got := pkg.GetName()

	if want != got {
		t.Errorf("\nGetName\n  want: %s\n  got.: %s\n", want, got)
	}
}