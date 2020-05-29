package wipchat

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrTokenRequired = Error("token required")
)
