package response

type format struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func New() format {
	return format{}
}
