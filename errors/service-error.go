package errors

// ServiceError is used to return error message
type ServiceError struct {
	Message string `json:"message"`
}
