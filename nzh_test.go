package gonzh

import "testing"

func TestSmallIntegerLower(t *testing.T) {
	want := "一千三百五十六"
	got := EncodeString("1356", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSmallMoneyCapital(t *testing.T) {
	want := "壹仟叁佰伍拾陆圆整"
	got := EncodeString("1356", true, true)
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

func TestBigIntegerCapital(t *testing.T) {
	want := "伍仟肆佰叁拾贰万壹仟玖佰伍拾捌"
	got := EncodeString("54321958", false, true)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestBigMoneyCapital(t *testing.T) {
	want := "伍仟肆佰叁拾贰万壹仟玖佰伍拾捌圆整"
	got := EncodeString("54321958", true, true)
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

func TestHugeMoneyCapital(t *testing.T) {
	want := "壹拾贰垓叁仟肆佰伍拾陆京柒仟捌佰玖拾捌兆柒仟陆佰伍拾肆亿叁仟贰佰壹拾贰万叁仟肆佰伍拾陆圆整"
	got := EncodeString("1234567898765432123456", true, true)
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

func TestExtremeIntegerZeros(t *testing.T) {
	want := "九十兆"
	got := EncodeString("90000000000000", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNegativeIntegerLower(t *testing.T) {
	want := "负八千九百七十六"
	got := EncodeString("-8976", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNegativeMoneyCapital(t *testing.T) {
	want := "负捌仟玖佰柒拾陆圆整"
	got := EncodeString("-8976", true, true)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSmallDecimalLower(t *testing.T) {
	want := "三点一四一五九二六"
	got := EncodeString("3.1415926", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestDecimalMoneyCapital(t *testing.T) {
	want := "叁圆壹角肆分"
	got := EncodeString("3.1415926", true, true)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestDecimalMoneyLower2(t *testing.T) {
	want := "一十五万元一角六分"
	got := EncodeString("150000.16", true, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNegativeDecimalLower(t *testing.T) {
	want := "负四百六十七万一千三百五十六点五四八"
	got := EncodeString("-4671356.548", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNegativeDecimalMoney(t *testing.T) {
	want := "负肆佰陆拾柒万壹仟叁佰伍拾陆圆伍角"
	got := EncodeString("-4671356.5", true, true)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNumberWithZeros(t *testing.T) {
	want := "负四百六十万零三百五十六点五四八"
	got := EncodeString("-4600356.548", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNumberZeroTail(t *testing.T) {
	want := "负五百六十"
	got := EncodeString("-560", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNumberMultipleZeros(t *testing.T) {
	want := "五十万零三百零六点零四"
	got := EncodeString("500306.04", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestNumberInterleavedZeros(t *testing.T) {
	want := "九千零八十万七千零六十点一零二零三"
	got := EncodeString("090807060.102030", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSciNotationPositive(t *testing.T) {
	want := "一十二垓三千四百五十六京七千八百九十兆"
	got := EncodeString("1.23456789e+21", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSciNotationNegative(t *testing.T) {
	want := "零点零零零零零零零零零三一四一五九二六"
	got := EncodeString("3.1415926e-10", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSciNotation3(t *testing.T) {
	want := "一千零三十四点五六"
	got := EncodeString("1.03456e3", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestSciNotation4(t *testing.T) {
	want := "一十二点三四五"
	got := EncodeString("1234.5e-2", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestIllegalNumber(t *testing.T) {
	want := "123af4.52"
	got := EncodeString("123af4.52", false, false)
	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}
