package reply

const (
	ERROR   = 7
	SUCCESS = 0
)

type RpcReply struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Reply(code int, data interface{}, msg string) RpcReply {
	return RpcReply{
		Code: code,
		Data: data,
		Message: msg,
	}
}

func Ok() RpcReply {
	return RpcReply{
		Code: SUCCESS,
		Message: "ok",
	}
}

func OkWithData(data interface{}) RpcReply {
	return RpcReply{
		Code: SUCCESS,
		Data: data,
		Message: "ok",
	}
}

func OkWithDetail(data interface{}, msg string) RpcReply {
	return RpcReply{
		Code: SUCCESS,
		Data: data,
		Message: msg,
	}
}

func Fail() RpcReply {
	return RpcReply{
		Code: ERROR,
		Message: "fail",
	}
}

func FailWithData(data interface{}) RpcReply {
	return RpcReply{
		Code: ERROR,
		Data: data,
		Message: "fail",
	}
}

func FailWithDetail(data interface{}, msg string) RpcReply {
	return RpcReply{
		Code: ERROR,
		Data: data,
		Message: msg,
	}
}