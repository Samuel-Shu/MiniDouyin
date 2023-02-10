package utils

const(
	SUCCESS = 0
	FAIL = 1
	USER_SUCCESS_REGISTER = 2100	//USER 类 2xxx
	USER_SUCCESS_LOGIN = 2101
	USER_FAIL_REGISTER = 2000
	USER_FAIL_LOGIN = 2001

)

var StatusMsg = map[int]string{
	SUCCESS: "成功",
	FAIL: "失败",
	USER_SUCCESS_REGISTER: "用户注册成功",
	USER_SUCCESS_LOGIN: "用户登录成功",
	USER_FAIL_REGISTER: "用户注册失败",
	USER_FAIL_LOGIN: "用户登录失败",
}

func GetStatusMsg(code int) string {
	return StatusMsg[code]
}