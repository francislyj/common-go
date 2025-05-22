package response

import (
	"github.com/francislyj/common-go/pkg/errorcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    errorcode.Success,
		Message: errorcode.Msg[errorcode.Success],
		Data:    data,
	})
}

// Error 支持只传 code，自动带默认 message，也支持传自定义 message
func Error(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	} else {
		m, ok := errorcode.Msg[code]
		if ok {
			msg = m
		} else {
			msg = "未知错误"
		}
	}
	c.JSON(200, Response{
		Code:    code,
		Message: msg,
	})
}
