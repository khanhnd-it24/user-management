package enums

type Role int

const (
	SuperAdmin Role = 1
	User       Role = 2
)

func (role Role) String() string {
	switch role {
	case SuperAdmin:
		return "SuperAdmin"
	case User:
		return "User"
	default:
		return "Unknown"
	}
}
