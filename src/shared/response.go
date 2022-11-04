package shared

type BaseResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MetaResponses struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
}

type BaseResponses struct {
	Success bool          `json:"success"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Meta    MetaResponses `json:"meta"`
	Data    []interface{} `json:"data"`
}
