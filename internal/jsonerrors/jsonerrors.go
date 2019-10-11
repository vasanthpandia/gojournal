
package jsonerrors

type Error struct {
	Message string `json:"error"`
}

func New(msg string) Error {
	jerr := Error {
		Message: msg,
	}
	return jerr
}

var ResourceNotFound Error = Error {
	Message: "Resource Not Found",
}

var BadRequest Error = Error {
	Message: "Bad Request",
}
