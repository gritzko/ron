package ron

func (vec VVector) Add(uuid UUID) {
	val, ok := vec[uuid[1]]
	if !ok || val < uuid[0] {
		vec[uuid[1]] = uuid[0]
	}
}

func (vec VVector) AddString(uuidString string) error {
	uuid, err := ParseUUIDString(uuidString)
	if err == nil {
		vec.Add(uuid)
	}
	return err
}

func (vec VVector) Get(uuid UUID) uint64 {
	return vec[uuid[1]]
}

func (vec VVector) GetUUID(uuid UUID) UUID {
	return UUID{vec[uuid[1]], uuid[1]}
}
