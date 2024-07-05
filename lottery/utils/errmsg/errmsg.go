package errmsg



const (
	ERROR = 500
	SUCCESS = 200
)



var codeMsg = map[int]string {
	ERROR: "FAIL",
	SUCCESS: "OK",
}


func GetErrMsg(code int) string {
	return codeMsg[code]
}