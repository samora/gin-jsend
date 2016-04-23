package jsend

import (
	"github.com/gin-gonic/gin"
)

// OK - 200 http status
func OK(c *gin.Context, data interface{}) {
	Success(c, 200, data)
}

// Created - 201 http status
func Created(c *gin.Context, data interface{}) {
	Success(c, 201, data)
}

// BadRequest - 400 http status
func BadRequest(c *gin.Context, data interface{}) {
	Fail(c, 400, data)
}

// Unauthorized - 401 http status
func Unauthorized(c *gin.Context, code int64, data interface{}) {
	Error(c, 401, "Unauthorized", code, data)
}

// Forbidden - 403 http status
func Forbidden(c *gin.Context, code int64, data interface{}) {
	Error(c, 403, "Forbidden", code, data)
}

// NotFound - 404 http status
func NotFound(c *gin.Context, code int64, data interface{}) {
	Error(c, 404, "Not found", code, data)
}

// TooManyRequests - 429 http status
func TooManyRequests(c *gin.Context, code int64, data interface{}) {
	Error(c, 429, "Too many requests", code, data)
}

// InternalServerError - 500 http status
func InternalServerError(c *gin.Context, msg string, code int64, data interface{}) {
	Error(c, 500, msg, code, data)
}
