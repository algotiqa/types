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

func TestNewSession(t *testing.T) {
	ts := TradingSession{
		Slots: []*TradingSlot{
			{Day: 0, Open: Time(1700), Close: Time(1600), EndSession: true},  //Sun 2026-03-15
			{Day: 1, Open: Time(1700), Close: Time(1600), EndSession: true},  //Mon 2026-03-16
			{Day: 2, Open: Time(1700), Close: Time(1000), EndSession: false}, //Tue 2026-03-17
			{Day: 3, Open: Time(1100), Close: Time(1500), EndSession: true},  //Wed 2026-03-18
			{Day: 4, Open: Time(1745), Close: Time(1950), EndSession: true},  //Thu 2026-03-19
		},
	}

	//--- No hole at all

	t1 := date(2026, 3, 15, 19, 03) //Sun
	t2 := date(2026, 3, 15, 21, 59)

	if ts.IsNewSession(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Real session end

	t1 = date(2026, 3, 16, 16, 00) //Mon
	t2 = date(2026, 3, 16, 17, 01)

	if !ts.IsNewSession(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Pause inside the session

	t1 = date(2026, 3, 18, 9, 15) //Wed
	t2 = date(2026, 3, 18, 12, 5)

	if ts.IsNewSession(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}

	//--- Week break

	t1 = date(2026, 3, 20, 19, 50) //Fri
	t2 = date(2026, 3, 22, 17, 5)

	if ts.IsNewSession(t1, t2) {
		t.Errorf("NewSession failed. Dates %v and %v", t1, t2)
	}
}

//=============================================================================

func date(y, m, d, hh, mm int) time.Time {
	return time.Date(y, time.Month(m), d, hh, mm, 00, 0, time.UTC)
}

//=============================================================================
