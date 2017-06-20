package gonzh

import "testing"

func TestSmallIntegerLower(t *testing.T) {
	want := "一千三百五十六"
	got := EncodeString("1356", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestBigIntegerLower(t *testing.T) {
	want := "五千四百三十二万一千九百五十八"
	got := EncodeString("54321958", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestHugeIntegerLower(t *testing.T) {
	want := "一十二垓三千四百五十六京七千八百九十八兆七千六百五十四亿三千二百一十二万三千四百五十六"
	got := EncodeString("1234567898765432123456", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestExtremeIntegerLower(t *testing.T) {
	want := "一十二载三千四百五十六正七千八百九十八涧七千六百五十四沟三千二百一十二穰三千四百五十六秭七千八百九十八垓七千六百五十四京三千二百一十二兆三千四百五十六亿七千八百九十八万七千六百五十四"
	got := EncodeString("1234567898765432123456789876543212345678987654", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSmallDecimalLower(t *testing.T) {
	want := "一千三百五十六点五四六八"
	got := EncodeString("1356.5468", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}
