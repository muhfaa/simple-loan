package common

import (
	"loanapp/business"
	"net/http"
)

const (
	errDuplicateData       = "duplicate_data"
	errInvalidSpec         = "invalid_spec"
	errDataNotFound        = "data_not_found"
	errInternalServerError = "internal_server_error"
)

//NewBusinessErrorMappingResponse Response return choosen http status like 200 OK, 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewBusinessErrorMappingResponse(err error) (int, DefaultResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, DefaultResponse) {
	switch err {
	default:
		return http.StatusInternalServerError, newInternalServerErrorResponse()
	case business.ErrNotFound:
		return http.StatusNotFound, newNotFoundResponse()
	case business.ErrDuplicate:
		return http.StatusBadRequest, newDuplicateResponse()
	case business.ErrInvalidSpec:
		return http.StatusBadRequest, newValidationResponse(err.Error())
	}
}

//newDuplicateResponse Duplicate Data format response
func newDuplicateResponse() DefaultResponse {
	return DefaultResponse{
		errDuplicateData,
		"Duplicate data",
		map[string]interface{}{},
	}
}

//newValidationResponse failed to validate request payload
func newValidationResponse(message string) DefaultResponse {
	return DefaultResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		errDataNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}
