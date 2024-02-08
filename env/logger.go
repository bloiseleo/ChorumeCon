package env

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ChorumeconLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[CHORUMECON] - %s - [%s] %s %s %s %d\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode, // Este corresponde ao %d na string de formatação
	)

}
