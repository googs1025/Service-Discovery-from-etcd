package api

import "fmt"

type ClientResult struct {
	Code int
	Msg string
}

var (
	EtcdError = &ClientResult{Code: 10001, Msg: "the ETCD connection error!"}
	RPCResponseError = &ClientResult{Code: 10002, Msg: "the grpc response error!"}
)


func (r ClientResult) Error() string {
	return fmt.Sprintf("Error[code=%d, msg=%s]", r.Code, r.Msg)
}

func (r ClientResult) FindMsg(msg string) *ClientResult {
	r.Msg = msg
	return &r
}

func (r ClientResult) FindCode(code int) *ClientResult {
	r.Code = code
	return &r
}
