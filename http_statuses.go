package jsend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// OK - 200 http status
func OK(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, data)
}

// Created - 201 http status
func Created(c *gin.Context, data interface{}) {
	Success(c, http.StatusCreated, data)
}

// BadRequest - 400 http status
func BadRequest(c *gin.Context, data interface{}) {
	Fail(c, http.StatusBadRequest, data)
}

// Unauthorized - 401 http status
func Unauthorized(c *gin.Context, code int64, data interface{}) {
	Error(c, http.StatusUnauthorized, "Unauthorized", code, data)
}

// Forbidden - 403 http status
func Forbidden(c *gin.Context, code int64, data interface{}) {
	Error(c, http.StatusForbidden, "Forbidden", code, data)
}

// NotFound - 404 http status
func NotFound(c *gin.Context, code int64, data interface{}) {
	Error(c, http.StatusNotFound, "Not found", code, data)
}

// InternalServerError - 500 http status
func InternalServerError(c *gin.Context, msg string, code int64, data interface{}) {
	Error(c, http.StatusInternalServerError, msg, code, data)
}
