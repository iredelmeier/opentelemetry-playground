package file

import "os"

type ErrorHandler interface {
	Handle(err error)
}

type DefaultErrorHandler struct {
	file *os.File
}

func (eh DefaultErrorHandler) Handle(err error) {
	_, _ = eh.file.WriteString(err.Error())
}
