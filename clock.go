package ron

import "time"
import (
	"encoding/binary"
	"encoding/hex"
)

// hybrid calendar/logical clock
type Clock struct {
	offset    time.Duration
	lastSeen  UUID
	Mode      UUID
	MinLength int
}

var CLOCK_CALENDAR = NewName("Calendar")
var CLOCK_EPOCH = NewName("Epoch") // TODO implement behavior
var CLOCK_LAMPORT = NewName("Logical")

var MAX_BIT_GRAB uint64 = 1 << 20

func NewClock(replica uint64, mode UUID, minLen int) Clock {
	origin := (replica & INT60_FULL) | UUID_UPPER_BITS[UUID_EVENT]
	return Clock{lastSeen: NewEventUUID(0, origin), Mode: mode, MinLength: minLen}
}

func EncodeCalendar(t time.Time) (i uint64) {
	months := (t.Year()-2010)*12 + int(t.Month()) - 1
	i |= uint64(months)
	days := t.Day() - 1
	i <<= 6
	i |= uint64(days)
	hours := t.Hour()
	i <<= 6
	i |= uint64(hours)
	minutes := t.Minute()
	i <<= 6
	i |= uint64(minutes)
	seconds := t.Second()
	i <<= 6
	i |= uint64(seconds)
	micros := t.Nanosecond() / 100
	i <<= 24
	i |= uint64(micros)
	return i
}

func CalendarToRFC(uuid UUID) (u [16]byte) {
	// the formula comes from satori/go.uuid
	time := DecodeCalendar(uuid.Value())
	timeRfc := uint64(122192928000000000) + uint64(time.UnixNano()/100)
	binary.BigEndian.PutUint32(u[0:], uint32(timeRfc))
	binary.BigEndian.PutUint16(u[4:], uint16(timeRfc>>32))
	binary.BigEndian.PutUint16(u[6:], uint16(timeRfc>>48))
	binary.BigEndian.PutUint16(u[8:], 0)

	var replicaId [6]byte
	orig := uuid.Origin()
	orig >>= 60 - 6*8
	for i := 0; i < 6; i++ {
		replicaId[5-i] = byte(orig & 255)
		orig >>= 8
	}

	copy(u[10:], replicaId[:])
	var version byte = 1
	u[6] = (u[6] & 0x0f) | (version << 4)
	u[8] = (u[8] & 0xbf) | 0x80

	return u
}

func CalendarToRFCString(uuid UUID) string {
	u := CalendarToRFC(uuid)
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)

}

func trim_time(full, last uint64) uint64 {
	i := 5
	for i < 11 && full&PREFIX_MASKS[i] <= last {
		i++
	}
	return full & PREFIX_MASKS[i]
}

func (clock *Clock) Time() UUID {
	var val uint64
	last := clock.lastSeen.Value()
	switch clock.Mode {
	case CLOCK_CALENDAR:
		t := time.Now().Add(clock.offset).UTC()
		val = EncodeCalendar(t)
	case CLOCK_LAMPORT:
		val = last + 1
	case CLOCK_EPOCH:
		t := time.Now().Add(clock.offset).UTC()
		val = uint64(t.Unix()) << (4 * 6) // TODO define
	}
	if val <= last {
		val = last + 1
	} else {
		val = trim_time(val, last)
	}
	ret := NewEventUUID(val, clock.lastSeen.Origin())
	clock.See(ret)
	return ret
}

func (clock *Clock) See(uuid UUID) bool {
	if clock.lastSeen.Value() < uuid.Value() {
		clock.lastSeen = NewEventUUID(uuid.Value(), clock.lastSeen.Origin())
	}
	return true
}

func (clock Clock) IsSane(uuid UUID) bool {
	switch clock.Mode {
	case CLOCK_LAMPORT:
		return clock.lastSeen.Value()+MAX_BIT_GRAB > uuid.Value()
	default:
		return true
	}
}

func (clock Clock) Decode(uuid UUID) time.Time {
	switch clock.Mode {
	case CLOCK_CALENDAR:
		return DecodeCalendar(uuid.Value())
	default:
		return time.Time{}
	}

}

const MASK24 uint64 = 16777215

func DecodeCalendar(v uint64) time.Time {
	var ns100 int = int(v & MASK24)
	v >>= 24
	var secs int = int(v & 63)
	v >>= 6
	var mins int = int(v & 63)
	v >>= 6
	var hours int = int(v & 63)
	v >>= 6
	var days int = int(v & 63)
	v >>= 6
	var months int = int(v & 4095)
	var month = months % 12
	var year = months / 12
	t := time.Date(year+2010, time.Month(month+1), days+1, hours, mins, secs, ns100*100, time.UTC)
	return t
}
