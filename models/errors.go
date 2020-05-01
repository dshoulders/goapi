package models

type ErrorType string

type TypedError struct {
	Type ErrorType
	Err  error
}

func (e *TypedError) Error() string { return e.Err.Error() }

type QueryError struct {
	Err error
}

func (e *QueryError) Error() string { return e.Err.Error() }

type NotFoundError struct {
	Err error
}

func (e *NotFoundError) Error() string { return e.Err.Error() }

type AuthenticationError struct {
	Err error
}

func (e *AuthenticationError) Error() string { return e.Err.Error() }
