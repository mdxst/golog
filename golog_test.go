package golog

import (
	"testing"
)

func TestLog(t *testing.T) {
	log := Log("ciao", "test", 0)
	t.Logf("output di Log(\"ciao\", \"test\", 0) era %v", log)
//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
}

func TestLogVuoto(t *testing.T) {
//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := Log("", "", 0)
	t.Logf("output di Log(\"\", \"\", 0) era %v", log)
}
