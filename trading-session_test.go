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
	"time"
)

//=============================================================================

var ts = &TradingSession{
	Slots: []*TradingSlot{
		{Day: 0, Open: Time(1700), Close: Time(1600), EndSession: true},  //Sun 2026-03-15
		{Day: 1, Open: Time(1700), Close: Time(1600), EndSession: true},  //Mon 2026-03-16
		{Day: 2, Open: Time(1700), Close: Time(1000), EndSession: false}, //Tue 2026-03-17
		{Day: 3, Open: Time(1100), Close: Time(1500), EndSession: true},  //Wed 2026-03-18
		{Day: 4, Open: Time(1745), Close: Time(1950), EndSession: true},  //Thu 2026-03-19
	},
}

var ts60m = &TradingSession{
	Slots: []*TradingSlot{
		{Day: 0, Open: Time(1700), Close: Time(1600), EndSession: true},  //Sun 2026-03-15
		{Day: 1, Open: Time(1700), Close: Time(1600), EndSession: true},  //Mon 2026-03-16
		{Day: 2, Open: Time(1700), Close: Time(1000), EndSession: false}, //Tue 2026-03-17
	},
}

var ts15m = &TradingSession{
	Slots: []*TradingSlot{
		{Day: 0, Open: Time(1715), Close: Time(1645), EndSession: true}, //Sun 2026-03-15
		{Day: 1, Open: Time(1745), Close: Time(1600), EndSession: true}, //Mon 2026-03-16
	},
}

var ts1m = &TradingSession{
	Slots: []*TradingSlot{
		{Day: 0, Open: Time(1716), Close: Time(1645), EndSession: true}, //Sun 2026-03-15
	},
}

var tsAnomalous = &TradingSession{
	Slots: []*TradingSlot{
		{Day: 0, Open: Time(1700), Close: Time(1600), EndSession: true}, //Sun 2026-03-15
		{Day: 2, Open: Time(2300), Close: Time(1600), EndSession: true}, //Mon 2026-03-16
	},
}

//=============================================================================

func TestNewSession(t *testing.T) {

	//--- No hole at all

	t1 := date(2026, 3, 15, 19, 03) //Sun
	t2 := date(2026, 3, 15, 21, 59)

	if ts.CrossSessions(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Real session end

	t1 = date(2026, 3, 16, 16, 00) //Mon
	t2 = date(2026, 3, 16, 17, 01)

	if !ts.CrossSessions(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Pause inside the session

	t1 = date(2026, 3, 18, 9, 15) //Wed
	t2 = date(2026, 3, 18, 12, 5)

	if ts.CrossSessions(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Week break

	t1 = date(2026, 3, 20, 19, 50) //Fri
	t2 = date(2026, 3, 22, 17, 5)

	if ts.CrossSessions(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Real session end (anomalous)

	t1 = date(2026, 3, 16, 16, 00) //Mon
	t2 = date(2026, 3, 16, 17, 01)

	if !tsAnomalous.CrossSessions(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}
}

//=============================================================================

func TestGranularity(t *testing.T) {
	granularity(t, ts, 5)
	granularity(t, ts60m, 60)
	granularity(t, ts15m, 15)
	granularity(t, ts1m, 1)
}

//=============================================================================

func TestTradingSlot(t *testing.T) {
	d    := date(2026, 3, 23, 12, 0) //--- Monday
	slot := ts.FindSlot(d)
	if slot == nil {
		t.Errorf("TradingSlot not found")
	} else if slot.Day != 0 {
		t.Errorf("Bad tradingslot day: expected %v but got %v", d, slot.Day)
	}

	d = date(2026, 3, 23, 19, 0) //--- Monday
	slot = ts.FindSlot(d)
	if slot == nil {
		t.Errorf("TradingSlot not found")
	} else if slot.Day != 1 {
		t.Errorf("Bad tradingslot day: expected %v but got %v", d, slot.Day)
	}

	d = date(2026, 3, 23, 16, 30) //--- Monday
	slot = ts.FindSlot(d)
	if slot != nil {
		t.Errorf("TradingSlot should have been null")
	}
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func date(y, m, d, hh, mm int) time.Time {
	return time.Date(y, time.Month(m), d, hh, mm, 00, 0, time.UTC)
}

//=============================================================================

func granularity(t *testing.T, ts *TradingSession, expected int) {
	if ts.Granularity() != expected {
		t.Errorf("Granularity failed. Expected %v", expected)
	}
}

//=============================================================================
