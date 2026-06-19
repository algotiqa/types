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
	"errors"
	"fmt"
	"strconv"
	"time"
)

//=============================================================================

type Date int

//=============================================================================

func NewDate(year int, month int, day int) Date {
	return Date(year*10000 + month*100 + day)
}

//=============================================================================

func (dt Date) Year() int {
	return int(dt / 10000)
}

//=============================================================================

func (dt Date) Month() int {
	return int((dt / 100) % 100)
}

//=============================================================================

func (dt Date) Day() int {
	return int(dt % 100)
}

//=============================================================================

func (dt Date) String() string {
	return fmt.Sprintf("%4d-%02d-%02d", dt.Year(), dt.Month(), dt.Day())
}

//=============================================================================

func (dt Date) IsNil() bool {
	return dt == 0
}

//=============================================================================

func (dt Date) IsValid() bool {
	if dt < 0 {
		return false
	}

	d := dt.Day()
	m := dt.Month()

	if m < 1 || m > 12 {
		return false
	}

	if m == 4 || m == 6 || m == 9 || m == 11 {
		return d >= 1 && d <= 30
	}

	if m == 2 {
		return d >= 1 && d <= 29
	}

	return d >= 1 && d <= 31
}

//=============================================================================

func (dt Date) ToDateTime(endDay bool, loc *time.Location) time.Time {
	hh := 0
	mm := 0
	ss := 0

	if endDay {
		hh = 23
		mm = 59
		ss = 59
	}

	return time.Date(dt.Year(), time.Month(dt.Month()), dt.Day(), hh, mm, ss, 0, loc)
}

//=============================================================================

func (dt Date) AddDays(days int) Date {
	t := dt.ToDateTime(false, time.UTC)
	t = t.Add(time.Duration(days) * 24 * time.Hour)

	y, m, d := t.Date()

	return Date(y*10000 + int(m)*100 + d)
}

//=============================================================================

func (dt Date) IsToday(loc *time.Location) bool {
	return dt == Today(loc)
}

//=============================================================================

func (dt Date) Days(d Date) int {
	ts := dt.ToDateTime(false, time.UTC)
	td := d.ToDateTime(false, time.UTC)
	seconds := td.Unix() - ts.Unix()

	return int(seconds / 3600 / 24)
}

//=============================================================================
//===
//=== General functions
//===
//=============================================================================

func ToDate(t *time.Time) Date {
	if t == nil {
		return 0
	}

	y, m, d := t.Date()

	return Date(y*10000 + int(m)*100 + d)
}

//=============================================================================

func ParseIntDate(value string, required bool) (Date, error) {
	if value == "" {
		if required {
			return 0, errors.New("value is required")
		}

		return 0, nil
	}

	d, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	id := Date(d)

	if !id.IsValid() {
		return 0, errors.New("invalid date")
	}

	return id, nil
}

//=============================================================================

func Today(loc *time.Location) Date {
	tn := time.Now().In(loc)
	return ToDate(&tn)
}

//=============================================================================
