package api

import "fmt"

type ServerResult struct {
	Code int
	Msg string
}

var (
	EtcdConnectionError = &ServerResult{Code: 11001, Msg: "the ETCD connection error!"}
	EtcdAddOrDeleteEndpointError = &ServerResult{Code: 11002, Msg: "the add or delete addr operation error!"}
	EtcdLeaseError = &ServerResult{Code: 11003, Msg: "the lease operation error!"}
	EtcdKeepAliveError = &ServerResult{Code: 11004, Msg: "the keep alive operation error!"}

)


func (r ServerResult) Error() string {
	return fmt.Sprintf("Error[code=%d, msg=%s]", r.Code, r.Msg)
}

func (r ServerResult) FindMsg(msg string) *ServerResult {
	r.Msg = msg
	return &r
}

func (r ServerResult) FindCode(code int) *ServerResult {
	r.Code = code
	return &r
}

