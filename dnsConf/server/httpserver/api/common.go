package api

type Request struct {
}

const (
	RespSuccess = iota
	RespErr
	UnmarshalSigContentErr
	SaveErr
)

type Response struct {
	Code   int64       `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Result interface{} `json:"result"`
}
