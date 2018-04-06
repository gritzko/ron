package ron

import "testing"

func TestUUID_Parse(t *testing.T) {
	samples := map[string]UUID{
		"1":                NewRonUUID(UUID_NAME, UUID_NAME_TRANSCENDENT, 1<<54, 0),
		"1/978$1400075997": NewRonUUID(UUID_NAME, UUID_NAME_ISBN, 164135095794401280, 19140298535113287),
	}
	for uuidStr, uuid := range samples {
		parsed, err := ParseUUIDString(uuidStr)
		if err != nil {
			t.Fail()
			t.Log("parse fail", err)
			break
		}
		if uuid != parsed {
			t.Fail()
			t.Logf("got %s expected %s", parsed.String(), uuid.String())
			break
		}
		if parsed.String() != uuidStr {
			t.Fail()
			t.Logf("serialized as %s expected %s", parsed.String(), uuidStr)
		}
	}

}
