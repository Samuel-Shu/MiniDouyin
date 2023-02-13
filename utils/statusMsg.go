package utils

const(
	SUCCESS int32 = 0
	FAIL int32= 1
	USER_SUCCESS_REGISTER int32 = 2100	//USER 类 2xxx
	USER_SUCCESS_LOGIN int32 = 2101
	USER_FAIL_REGISTER int32 = 2000
	USER_FAIL_LOGIN int32 = 2001

	ERROR_TOKEN_EXIST     int32 = 1004
	ERROR_TOKEN_RUNTIME    int32 = 1005
	ERROR_TOKEN_WRONG     int32 = 1006
	ERROR_TOKEN_TYPE_WRONG int32 = 1007
)

var StatusMsg = map[int32]string{
	SUCCESS: "成功",
	FAIL: "失败",
	USER_SUCCESS_REGISTER: "用户注册成功",
	USER_SUCCESS_LOGIN: "用户登录成功",
	USER_FAIL_REGISTER: "用户注册失败",
	USER_FAIL_LOGIN: "用户登录失败",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
}

func GetStatusMsg(code int32) string {
	return StatusMsg[code]
}