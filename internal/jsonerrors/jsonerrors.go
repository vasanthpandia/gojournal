
package jsonerrors

type Error struct {
	message string `json:"error"`
}

func New(msg string) Error {
	return Error {
		message: msg,
	}
}
