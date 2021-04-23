package error

import "fmt"

type myError struct{
	fileName string
	line  int
	msg  string
}

type myErrorOptions func(*myError)

func FileName(fileName string) myErrorOptions {
	return func(m *myError) {
		m.fileName = fileName
	}
}

func Line(line int) myErrorOptions {
	return func(m *myError) {
		m.line = line
	}
}

func Msg(msg string) myErrorOptions {
	return func(m *myError) {
		m.msg = msg
	}
}

func NewMyError(opts...myErrorOptions) error {
	e:=&myError{}
	for _,o:=range opts{
		o(e)
	}
	return e
}

func (e *myError)Error() string  {
	return fmt.Sprintf("[%s:%d] %s",e.fileName,e.line,e.msg)
}

func init()  {
	fmt.Println("error init")
}