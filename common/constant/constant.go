package constant

const (
	Success = iota + 200
	DuplicatedUser
	UserNotFound
	WrongPassword

	InternalServerError = 500
)

func CodeInfo(constValue int) *string {
	constName := ""
	switch constValue {
	case Success:
		constName = "Success"
	case DuplicatedUser:
		constName = "DuplicatedUser"
	case UserNotFound:
		constName = "UserNotFound"
	case WrongPassword:
		constName = "WrongPassword"
	case InternalServerError:
		constName = "InternalServerError"
	default:
		constName = "Unknown"
	}
	return &constName
}
