package enums

type TypeSchedule string

const (
	Offline TypeSchedule = "offline"
)

func IsValidTypeSchedule(val TypeSchedule) bool {
	return val == Offline
}
