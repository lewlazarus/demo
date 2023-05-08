package msg

// Response represents response data
type Response struct {
	Res   float64 `json:"res"`
	Ratio float64 `json:"ratio"`
}

func NewResponse(res float64, ratio float64) *Response {
	return &Response{Res: res, Ratio: ratio}
}
