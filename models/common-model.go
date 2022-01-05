package models

type (
	CommonResponse struct {
		Msg     string
		Data    interface{}
		Latency float64
	}
)
