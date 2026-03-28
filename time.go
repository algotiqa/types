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
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//=============================================================================

const NilValue int = 9999

type Time int16

//=============================================================================

func (t Time) Hour() int {
	return int(t / 100)
}

//=============================================================================

func (t Time) Minute() int {
	return int(t % 100)
}

//=============================================================================

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

//=============================================================================

func (t Time) IsNil() bool {
	return t == 9999
}

//=============================================================================

func (t Time) IsValid() bool {
	if t < 0 {
		return false
	}

	h := t.Hour()
	m := t.Minute()

	if h < 0 || h > 23 {
		return false
	}
	if m < 0 || m > 59 {
		return false
	}

	return true
}

//=============================================================================

func (t Time) AsMinutes() int {
	hh := int(t / 100)
	mm := int(t % 100)

	return hh * 60 + mm
}

//=============================================================================

func (t Time) AddMinutes(mins int) Time {
	totMins := t.Hour()*60 + t.Minute()
	finMins := totMins + mins

	hh := finMins / 60
	mm := finMins % 60

	return NewTime(hh, mm)
}

//=============================================================================

func (t Time) Add(time Time) Time {
	h := t.Hour() + time.Hour()
	m := t.Minute() + time.Minute()

	if m >= 60 {
		m -= 60
		h++
	}

	if h >= 24 {
		h -= 24
	}

	return NewTime(h, m)
}

//=============================================================================

func (t Time) Sub(time Time) Time {
	h := t.Hour() - time.Hour()
	m := t.Minute() - time.Minute()

	if m < 0 {
		m += 60
		h--
	}

	if h < 0 {
		h += 24
	}

	return NewTime(h, m)
}

//=============================================================================
//===
//=== General functions
//===
//=============================================================================

func NewTime(hours, mins int) Time {
	return Time(hours*100 + mins)
}

//=============================================================================

func ParseIntTime(value string, required bool) (Time, error) {
	if value == "" {
		if required {
			return 0, errors.New("value is required")
		}

		return Time(NilValue), nil
	}

	t, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	it := Time(t)

	if !it.IsValid() {
		return 0, errors.New("invalid time")
	}

	return it, nil
}

//=============================================================================

func ParseStringTime(time string) (Time, error) {
	index := strings.Index(time, ":")
	if index < 1 || index > 2 || len(time) < 4 || len(time) > 5 {
		return 0, errors.New("bad time format")
	}

	hour, err1 := strconv.Atoi(time[0:index])
	mins, err2 := strconv.Atoi(time[index:])

	if err1 != nil {
		return 0, errors.New("time hour is not an integer")
	}

	if err2 != nil {
		return 0, errors.New("time minute is not an integer")
	}

	t := NewTime(hour, mins)
	if !t.IsValid() {
		return 0, errors.New("bad hour/minute value")
	}

	return t, nil
}

//=============================================================================
