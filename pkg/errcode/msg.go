package errcode

var msgMap = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	INVALID_PARAMS: "请求参数错误",
	RATE_LIMIT:     "访问限流",
	SYSTEM_ERROR:   "系统错误",
}

func GetMsg(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}

	return msgMap[ERROR]
}
