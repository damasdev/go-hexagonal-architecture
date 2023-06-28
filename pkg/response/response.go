package response

type defaultResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func NewErrResponse(
	code int,
	message string,
) defaultResponse {
	response := defaultResponse{}
	response.Status.Code = code
	response.Status.Message = message

	return response
}

func NewResponse(
	code int,
	message string,
	data any,
) defaultResponse {
	response := defaultResponse{}
	response.Status.Code = code
	response.Status.Message = message
	response.Data = data

	return response
}
