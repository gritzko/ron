package ron

type UUIDMultiMap struct {
	subs map[UUID][]UUID
	// avoid allocating lots of small slices, allocate larger slabs
	slab []UUID
}

// min RAM use per object: 16b key, 16b slice, 16*2b on the slab = 64

const U2M_DEFAULT_SLICE_SIZE = 2
const U2M_SLAB_SIZE = 32 // 32*16=512 bytes

func MakeUUID2Map() UUIDMultiMap {
	um := UUIDMultiMap{
		subs: make(map[UUID][]UUID),
		slab: make([]UUID, U2M_SLAB_SIZE),
	}
	return um
}

func (um *UUIDMultiMap) Add(key, value UUID) {
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

func (um UUIDMultiMap) Len() int {
	return len(um.subs)
}

func (um UUIDMultiMap) Keys() []UUID {
	ret := make([]UUID, 0, um.Len())
	for key, _ := range um.subs {
		ret = append(ret, key)
	}
	return ret
}

func (um UUIDMultiMap) List(key UUID) []UUID {
	return um.subs[key]
}

func (um UUIDMultiMap) Take(key UUID) (ret []UUID) {
	ret = um.List(key)
	delete(um.subs, key)
	return
}

func (um *UUIDMultiMap) Put(key UUID, values []UUID) {
	if len(values) > 0 {
		um.subs[key] = values
	} else {
		delete(um.subs, key)
	}
}

func (um *UUIDMultiMap) Remove(key UUID, value UUID) {
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
