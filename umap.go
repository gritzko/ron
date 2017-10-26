package ron

type UUID2Map struct {
	subs map[UUID][]UUID
	// avoid allocating lots of small slices, allocate larger slabs
	slab []UUID
}

// min RAM use per object: 16b key, 16b slice, 16*2b on the slab = 64

const U2M_DEFAULT_SLICE_SIZE = 2
const U2M_SLAB_SIZE = 32 // 32*16=512 bytes

func MakeUUID2Map() UUID2Map {
	um := UUID2Map{
		subs: make(map[UUID][]UUID),
		slab: make([]UUID, U2M_SLAB_SIZE),
	}
	return um
}

func (um *UUID2Map) Add(key, value UUID) {
	values, ok := um.subs[key]
	if !ok {
		if len(um.slab) < U2M_DEFAULT_SLICE_SIZE {
			um.slab = make([]UUID, U2M_SLAB_SIZE)
		}
		values = um.slab[0:0:U2M_DEFAULT_SLICE_SIZE]
		um.slab = um.slab[U2M_DEFAULT_SLICE_SIZE:]
	}
	values = append(values, value)
	um.Put(key, values)
}

func (um UUID2Map) List(key UUID) []UUID {
	return um.subs[key]
}

func (um *UUID2Map) Put(key UUID, values []UUID) {
	if len(values) > 0 {
		um.subs[key] = values
	} else {
		delete(um.subs, key)
	}
}

func (um *UUID2Map) Remove(key UUID, value UUID) {
	values, ok := um.subs[key]
	if !ok {
		return
	}
	for i := 0; i < len(values); i++ {
		if values[i] == value {
			l1 := len(values) - 1
			values[i] = values[l1]
			values = values[:l1]
			i--
		}
	}
	// TODO perf N^2
	// [ ] um.rm - list of removals for len(values)>100?
	// [ ] on every iteration, check against um.rm
	// [ ] um.rm over limit => iterate/merge
}
