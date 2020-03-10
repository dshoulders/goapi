package models

type Status string

const (
	Success Status = "success"
	Fail    Status = "fail"
	Error   Status = "error"
)
