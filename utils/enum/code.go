package enum

const (
	SUCCESS   = 200
	ERROR     = 500
	FORBIDDEN = 403

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED        = 1001
	ERROR_PASSWORD_WRONG       = 1002
	ERROR_USER_NOT_EXIST       = 1003
	ERROR_TOKEN_EXIST          = 1004
	ERROR_TOKEN_RUNTIME        = 1005
	ERROR_TOKEN_WRONG          = 1006
	ERROR_TOKEN_TYPE_WRONG     = 1007
	ERROR_USER_MIDDLEWARE_AUTH = 1008
	ERROR_CASBIN_LOCAL         = 1009
	ERROR_CASBIN_DYNAMICALLY   = 1010
	ERROR_USER_AUTHORIZING     = 1011
	ERROR_USER_NO_RIGHT        = 1012

	// code= 2000... 文章模块的错误

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:   "OK",
	ERROR:     "FAIL",
	FORBIDDEN: "FORBIDDEN",

	ERROR_USERNAME_USED:        "用户名已存在！",
	ERROR_PASSWORD_WRONG:       "密码错误",
	ERROR_USER_NOT_EXIST:       "用户不存在",
	ERROR_TOKEN_EXIST:          "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:        "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:          "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG:     "TOKEN格式错误,请重新登陆",
	ERROR_USER_MIDDLEWARE_AUTH: "用户登录授权失败",
	ERROR_CASBIN_LOCAL:         "未能创建 casbin 强制执行器",
	ERROR_CASBIN_DYNAMICALLY:   "无法从数据库加载策略",
	ERROR_USER_AUTHORIZING:     "授权用户时发生错误",
	ERROR_USER_NO_RIGHT:        "该用户无权限",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

// GetCodeMsg 获取信息
func GetCodeMsg(code int) string {
	return codeMsg[code]
}
