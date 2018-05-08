package ron

var CLOCK_CALENDAR UUID
var CLOCK_EPOCH UUID
var CLOCK_LAMPORT UUID

func init () {
	CLOCK_CALENDAR = NewName("Calendar")
	CLOCK_EPOCH = NewName("Epoch") // TODO implement behavior
	CLOCK_LAMPORT = NewName("Logical")
}
