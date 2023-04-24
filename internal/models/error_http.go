package models

type ErrorDetail struct {
	Detail string
}

type HTTPError struct {
	Message string
	Errors  []ErrorDetail
}
