package common

const (
	errBadRequest = "bad_request"
	errForbidden  = "forbidden"
)

type DefaultResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		errBadRequest,
		"Bad request",
		map[string]interface{}{},
	}
}

func NewInvalidRequestFormatResponse() DefaultResponse {
	return DefaultResponse{
		errBadRequest,
		"Invalid request format",
		map[string]interface{}{},
	}
}

func NewForbiddenResponse() DefaultResponse {
	return DefaultResponse{
		errForbidden,
		"Forbidden",
		map[string]interface{}{},
	}
}

func NewValidationResponse(message string) DefaultResponse {
	return DefaultResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}

func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		errDataNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

func NewInternalServerError(message string) DefaultResponse {
	return DefaultResponse{
		errInternalServerError,
		message,
		map[string]interface{}{},
	}
}
