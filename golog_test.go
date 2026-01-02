package golog

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	LivelloMax = 0
	t.Logf("impostato LivelloMax a 0.")
	log := Log(0, "rand!!", "test", ":)")
	t.Logf("output di Log(\"DEBUG\", \"test\", 0) era %v", log)
	// t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
}

func TestLogVuoto(t *testing.T) {
	LivelloMax = -1
	t.Logf("impostato LivelloMax a -1.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := Log(0, "", "")
	t.Logf("output di Log(\"\", \"\", 0) era %v", log)
}

func TestLogE(t *testing.T) {
	LivelloMax = 571
	t.Logf("impostato LivelloMax a 571.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogE(errors.New("test a a a errore finto!"))
	t.Logf("output di LogE(err) era %v", log)
}

func TestLogW(t *testing.T) {
	LivelloMax = 2
	t.Logf("impostato LivelloMax a 2.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogW("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogW(\"test\") era %v", log)
}

func TestLogI(t *testing.T) {
	LivelloMax = 1
	t.Logf("impostato LivelloMax a 1.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)

	log := LogI("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogI(\"test\") era %v", log)
}

func TestLogD(t *testing.T) {
	LivelloMax = 6
	t.Logf("impostato LivelloMax a 6.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogD("test", "test1", "dovrebbe", "essere", "staccato!")
	t.Logf("output di LogD(\"test\") era %v", log)
}

func TestLogS(t *testing.T) {
	LivelloMax = 213
	t.Logf("impostato LivelloMax a 213.")
	//t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	log := LogS("test", "test1", "titolo dovrebbe essere \"TestLogS\"!")
	t.Logf("output di LogS(\"test\") era %v", log)
}
