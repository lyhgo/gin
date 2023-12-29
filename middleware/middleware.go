package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetUP(engine *gin.Engine) {

	engine.Use(JsonLogger())
	engine.Use(ZapRecovery())
}
