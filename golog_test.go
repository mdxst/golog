package golog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	log := Log("DEBUG", "test", 0)
	t.Logf("output di Log(\"DEBUG\", \"test\", 0) era %v", log)
	// t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
}

func TestLogVuoto(t *testing.T) {
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := Log("", "", 0)
	t.Logf("output di Log(\"\", \"\", 0) era %v", log)
}

func TestLogE(t *testing.T) {
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogE(errors.New("test a a a errore finto!"))
	t.Logf("output di LogE(err) era %v", log)
}

func TestLogW(t *testing.T) {
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogW("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogW(\"test\") era %v", log)
}

func TestLogI(t *testing.T) {
	LivelloMax = 1
	t.Logf("impostato LivelloMax a 1.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)

	log := LogI("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogW(\"test\") era %v", log)
}

func TestLogD(t *testing.T) {
	LivelloMax = 3
	t.Logf("impostato LivelloMax a 3.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogD("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogW(\"test\") era %v", log)
}
