//=============================================================================
/*
Copyright © 2024 Andrea Carboni andrea.carboni71@gmail.com

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
	"encoding/json"
	"errors"
	"time"
)

//=============================================================================
//===
//=== TradingSession
//===
//=============================================================================

type TradingSession struct {
	Slots []*TradingSlot `json:"slots"`
}

//=============================================================================

func NewTradingSession(config string) (*TradingSession, error) {
	var sess TradingSession
	err := json.Unmarshal([]byte(config), &sess)
	if err != nil {
		return nil, errors.New("session is not a valid JDON: " + config)
	}

	return &sess, nil
}

//=============================================================================
//=== CrossSections time must be in data product's timezone

func (ts *TradingSession) CrossSessions(prev time.Time, next time.Time) bool {
	return ts.crossSlots(prev, next, true)
}

//=============================================================================
//=== CrossSlots time must be in data product's timezone

func (ts *TradingSession) CrossSlots(prev time.Time, next time.Time) bool {
	return ts.crossSlots(prev, next, false)
}

//=============================================================================

func (ts *TradingSession) Granularity() int {
	g05 := true
	g15 := true
	g60 := true

	for _, s := range ts.Slots {
		openMin := s.Open.Minute()
		closeMin := s.Close.Minute()

		if openMin != 0 || closeMin != 0 {
			g60 = false
		}

		if openMin%15 != 0 || closeMin%15 != 0 {
			g15 = false
		}

		if openMin%5 != 0 || closeMin%5 != 0 {
			g05 = false
		}
	}

	if g60 {
		return 60
	}

	if g15 {
		return 15
	}

	if g05 {
		return 5
	}

	return 1
}

//=============================================================================

func (ts *TradingSession) crossSlots(prev time.Time, next time.Time, endSession bool) bool {
	ph, pm, _ := prev.Clock()
	prevTime := NewTime(ph, pm)
	prevDow := int(prev.Weekday())
	nh, nm, _ := next.Clock()
	nextTime := NewTime(nh, nm)
	nextDow := int(next.Weekday())

	for _, s := range ts.Slots {
		if s.IsInside(prevDow, prevTime) && (!endSession || s.EndSession) {
			if !s.IsInside(nextDow, nextTime) {
				return true
			}
		}
	}

	return false
}

//=============================================================================

func (ts *TradingSession) FindSlot(t time.Time) *TradingSlot {
	ph, pm, _ := t.Clock()
	dayTime := NewTime(ph, pm)
	dayWeek := int(t.Weekday())

	for _, s := range ts.Slots {
		if s.IsInside(dayWeek, dayTime) {
			return s
		}
	}

	return nil
}

//=============================================================================
//===
//=== TradingSlot
//===
//=============================================================================

type TradingSlot struct {
	Day        int  `json:"day"`
	Open       Time `json:"open"`
	Close      Time `json:"close"`
	EndSession bool `json:"end"`
}

//=============================================================================

func (s *TradingSlot) IsInside(dow int, t Time) bool {
	if s.Open < s.Close {
		if s.Day == dow {
			return s.Open < t && t <= s.Close
		}
		return false
	}

	if s.Open > s.Close {
		return (s.Day == dow && s.Open < t) || (dow == s.Day+1 && t <= s.Close)
	}

	return false
}

//=============================================================================

func (s *TradingSlot) MinutesSinceOpen(t time.Time) int {
	ph, pm, _ := t.Clock()
	dayTime := NewTime(ph, pm)
	dayWeek := int(t.Weekday())

	mins := dayTime.AsMinutes() - s.Open.AsMinutes()

	if s.Day != dayWeek {
		mins += 1440
	}

	return mins
}

//=============================================================================
