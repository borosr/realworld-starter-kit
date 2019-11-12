package errors

type Error struct {
	Errors Errors `json:"errors"`
}

type Errors struct {
	Body []string `json:"body"`
}

func New(msgs []string) *Error {
	return &Error{
		Errors: Errors{msgs},
	}
}
