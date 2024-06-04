package msgcode


const (
	SUCCESS = 200
	ERROR = 500

	// 用户相关的错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
)

// 错误码对应的信息
var codeMsg = map[int]string{
	SUCCESS:				"OK",
	ERROR:					"FAIL",
	ERROR_USER_NOT_EXIST:	"用户不存在",
	ERROR_USERNAME_USED:	"用户已存在",
	ERROR_PASSWORD_WRONG:	"密码错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}