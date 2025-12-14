package golog

import "testing"

func TestLog(t *testing.T) {
	Log("ciao", "test", 0)
//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
}

func TestLogVuoto(t *testing.T) {
//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	Log("", "", 0)
}
