package errorcode

const (
	// 通用错误
	Success       = 0
	InvalidParam  = 400
	Unauthorized  = 401
	Forbidden     = 403
	NotFound      = 404
	InternalError = 500

	// 业务自定义错误
	UserNotFound  = 1001
	PasswordError = 1002
)

var Msg = map[int]string{
	Success:       "success",
	InvalidParam:  "参数错误",
	Unauthorized:  "未授权",
	Forbidden:     "无权限",
	NotFound:      "资源不存在",
	InternalError: "服务器内部错误",
	UserNotFound:  "用户不存在",
	PasswordError: "密码错误",
}
