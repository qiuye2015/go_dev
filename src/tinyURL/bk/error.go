package bk

type Error interface {
	error
	Status() int
}

type ErrorStatus struct {
	Code int
	Err  error
}

func (e ErrorStatus) Error() string {
	return e.Err.Error()
}

func (e ErrorStatus) Status() int {
	return e.Code
}
