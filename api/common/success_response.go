package common

const (
	Success = "success"
	Failed  = "failed"
)

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		string(Success),
		string(Success),
		data,
	}
}

//NewSuccessResponseNoData create new success payload
func NewSuccessResponseNoData() DefaultResponse {
	return DefaultResponse{
		string(Success),
		string(Success),
		map[string]interface{}{},
	}
}
