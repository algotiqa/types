//=============================================================================
//===
//=== Copyright (C) 2024-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
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
