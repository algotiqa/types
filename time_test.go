//=============================================================================
/*
Copyright © 2026 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
