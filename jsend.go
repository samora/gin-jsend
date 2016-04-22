package jsend

import (
	"github.com/gin-gonic/gin"
	"github.com/samora/jsend"
)

// Success JSend response
func Success(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(httpStatus, jsend.Success(data))
}

// Fail JSend response
func Fail(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(httpStatus, jsend.Fail(data))
}

// Error JSend response
func Error(c *gin.Context, httpStatus int, msg string, code int64, data interface{}) {
	c.JSON(httpStatus, jsend.Error(msg, code, data))
}
