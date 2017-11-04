package ron

import "testing"
import (
	"time"
)

func TestClock_Format(t *testing.T) {

	tests := [][]string{
		{"Fri Jan  1 00:00:00 UTC 2010", "0"},
		{"Sat May  1 01:02:00 UTC 2010", "04012"},
		{"Fri May 27 20:50:00 UTC 2016", "1CQKn"},
	}

	for i, pair := range tests {
		date, err := time.Parse(time.UnixDate, pair[0])
		if err != nil {
			t.Fail()
			break
		}
		ui := EncodeCalendar(date)
		s := FormatInt([]byte{}, ui)
		str := string(s)
		if str != pair[1] {
			t.Logf("case %d: %s must be %s", i, str, pair[1])
			t.Fail()
		}
		t2 := DecodeCalendar(ui)
		str2 := t2.Format(time.UnixDate)
		if str2 != pair[0] {
			t.Logf("case %d: %s must be %s", i, str2, pair[0])
			t.Fail()
		}
	}

}

func BenchmarkClock_Time(b *testing.B) {
	var prev UUID = ZERO_UUID
	var clock = Clock{}
	clock.lastSeen = NewEventUUID(0, 1)
	for i := 0; i < b.N; i++ {
		next := clock.Time()
		if next.Value() <= prev.Value() {
			b.Fail()
			b.Logf("%s (%d) <= %s (%d) at %d\n", next.String(), next.Value, prev.String(), prev.Value, i)
			break
		}
		prev = next
	}
}
