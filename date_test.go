//=============================================================================
//===
//=== Copyright (C) 2023-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package types

import (
	"testing"
	"time"
)

//=============================================================================

func TestToIntDate(t *testing.T) {
	td := time.Date(2025, 05, 03, 11, 12, 13, 0, time.UTC)
	id := ToDate(&td)
	exp := Date(20250503)

	if id != exp {
		t.Errorf("ToIntDate failed. Expected %v but got %v", exp, id)
	}
}

//=============================================================================

func TestStringDate(t *testing.T) {
	id := Date(20250503)
	sd := id.String()
	exp := "2025-05-03"

	if sd != exp {
		t.Errorf("String() failed. Expected %v but got %v", exp, sd)
	}
}

//=============================================================================

func TestParseIntDate(t *testing.T) {
	sd := "20250503"
	exp := Date(20250503)

	id, err := ParseIntDate(sd, true)
	if err != nil {
		t.Errorf("ParseIntDate failed. Expected %v but got %v", exp, id)
	}

	//---

	sd = "-20250503"

	id, err = ParseIntDate(sd, true)
	if err == nil {
		t.Errorf("ParseIntDate failed. Date is indicated as valid but it is not: %v", id)
	}

	//---

	sd = ""

	id, err = ParseIntDate(sd, false)
	if err != nil || !id.IsNil() {
		t.Errorf("ParseIntDate failed. Date is nil but got a valid date: %v", id)
	}
}

//=============================================================================

func TestDays(t *testing.T) {
	s := Date(20250503)
	d := Date(20250505)

	if s.Days(d) != 2 {
		t.Errorf("Days failed. Expected %v but got %v", 2, s.Days(d))
	}
}

//=============================================================================

func TestDaysLeap(t *testing.T) {
	s := Date(20240302)
	d := Date(20240228)

	if s.Days(d) != -3 {
		t.Errorf("Days failed. Expected %v but got %v", -3, s.Days(d))
	}
}

//=============================================================================
