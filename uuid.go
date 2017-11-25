package ron

func (uuid UUID) Value() uint64 {
	return uuid[0]
}

func (uuid UUID) Origin() uint64 {
	return uuid[1] & INT60_FULL
}

func Compare(a, b UUID) int {
	diff := int64(a.Value()) - int64(b.Value())
	if diff == 0 {
		diff = int64(a.Origin()) - int64(b.Origin())
	}
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func (t UUID) Equal(b UUID) bool {
	return t == b
}

func (uuid UUID) IsTranscendentName() bool {
	return uuid[1] == UUID_NAME_UPPER_BITS
}

func (uuid UUID) IsName() bool {
	return uuid.Scheme()==UUID_NAME
}

func (a UUID) LaterThan(b UUID) bool {
	if a.Value() == b.Value() {
		return a.Origin() > b.Origin()
	} else {
		return a.Value() > b.Value()
	}
}

func (a UUID) EarlierThan(b UUID) bool {
	// FIXME define through Compare
	if a.Value() == b.Value() {
		return a.Origin() < b.Origin()
	} else {
		return a.Value() < b.Value()
	}
}

func (a UUID) Scheme() uint64 {
	return a[1] >> 60
}

func (a UUID) Sign() byte {
	return UUID_PUNCT[uint(a.Scheme())]
}

func (a UUID) Replica() uint64 {
	return a[1] & INT60_FULL
}

func (a UUID) SameAs(b UUID) bool {
	if a.Value() != b.Value() {
		return false
	} else if a.Origin() == b.Origin() {
		return true
	} else if (a.Origin()^b.Origin())&INT60_FULL != 0 {
		return false
	} else {
		return a.Origin()&INT60_FULL == b.Origin()&INT60_FULL
	}
}

func (uuid UUID) Derived() UUID {
	if uuid.Scheme() == UUID_EVENT {
		return NewUUID(UUID_DERIVED, uuid.Value(), uuid.Origin())
	} else {
		return uuid
	}
}

var UUID_UPPER_BITS = [4]uint64{0, 1 << 60, 2 << 60, 3 << 60}

func NewUUID(scheme uint, time, origin uint64) UUID {
	return UUID{time, (origin & INT60_FULL) | UUID_UPPER_BITS[scheme]}
}

func NewEventUUID(time, origin uint64) UUID {
	return NewUUID(UUID_EVENT, time, origin)
}

func NewNameUUID(time, origin uint64) UUID {
	return NewUUID(UUID_NAME, time, origin)
}

func NewHashUUID(time, origin uint64) UUID {
	return NewUUID(UUID_HASH, time, origin)
}

// use for static strings only - panics on error
func NewName(name string) UUID {
	nam, err := ParseUUIDString(name)
	if err != nil {
		panic("bad name")
	}
	return nam
}

func NewError(name string) UUID {
	nam, err := ParseUUIDString(name)
	if err != nil {
		panic("bad error name")
	}
	return NewNameUUID(nam.Value(), INT60_ERROR)
}

func (uuid UUID) IsTemplate() bool {
	return uuid.Sign() == UUID_NAME && uuid.Value() == 0 && uuid.Origin() != 0
}

func (uuid UUID) ToScheme(scheme uint) UUID {
	return NewUUID(scheme, uuid.Value(), uuid.Origin())
}

func (uuid UUID) IsZero() bool {
	return uuid.Value() == 0 && uuid.Origin() == 0
}

func (uuid UUID) IsError() bool {
	return uuid.Origin() == INT60_ERROR
}

func (uuid UUID) ZipString(context UUID) string {
	var arr [INT60LEN*2 + 2]byte
	return string(FormatZipUUID(arr[:0], uuid, context))
}

func (uuid UUID) String() (ret string) {
	ret = uuid.ZipString(ZERO_UUID)
	if len(ret) == 0 {
		ret = "0"
	}
	return
}

func (uuid UUID) Error() string {
	if uuid.IsError() {
		return uuid.String()
	} else {
		return ""
	}
}
