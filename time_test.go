//=============================================================================
//===
//=== Copyright (C) 2026-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

import (
	"testing"
)

//=============================================================================

func TestStringTime(t *testing.T) {
	it := Time(905)
	st := it.String()
	exp := "09:05"

	if st != exp {
		t.Errorf("String failed. Expected %v but got %v", exp, st)
	}
}

//=============================================================================

func TestParseIntTime(t *testing.T) {
	st := "1245"
	exp := Time(1245)

	it, err := ParseIntTime(st, true)
	if err != nil {
		t.Errorf("ParseIntTime failed. Expected %v but got %v", exp, it)
	}

	//---

	st = "-234"

	it, err = ParseIntTime(st, true)
	if err == nil {
		t.Errorf("ParseIntTime failed. Time is indicated as valid but it is not: %v", it)
	}

	//---

	st = ""

	it, err = ParseIntTime(st, false)
	if err != nil || !it.IsNil() {
		t.Errorf("ParseIntTime failed. Time is nil but got a valid date: %v", it)
	}
}

//=============================================================================

func TestAddMinutes(t *testing.T) {
	it := Time(1025).AddMinutes(30)
	exp := Time(1055)

	if it != exp {
		t.Errorf("AddMinutes failed. Expected %v but got %v", exp, it)
	}

	it = Time(1025).AddMinutes(45)
	exp = Time(1110)

	if it != exp {
		t.Errorf("AddMinutes failed. Expected %v but got %v", exp, it)
	}

	it = Time(1025).AddMinutes(-30)
	exp = Time(955)

	if it != exp {
		t.Errorf("AddMinutes failed. Expected %v but got %v", exp, it)
	}

	it = Time(1025).AddMinutes(-61)
	exp = Time(924)

	if it != exp {
		t.Errorf("AddMinutes failed. Expected %v but got %v", exp, it)
	}
}

//=============================================================================

func TestBeforeTime(t *testing.T) {
	t1 := Time(1025)
	t2 := Time(1055)

	if !(t1 < t2) {
		t.Errorf("Before time failed. Dates %v and %v", t1, t2)
	}

	t1 = Time(1159)
	t2 = Time(921)

	if t1 < t2 {
		t.Errorf("Before time failed. Dates %v and %v", t1, t2)
	}

	t1 = Time(1234)

	if t1 < t1 {
		t.Errorf("Before time failed. Dates %v and %v", t1, t1)
	}
}

//=============================================================================
