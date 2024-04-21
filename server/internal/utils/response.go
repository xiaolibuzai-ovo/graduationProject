package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespError struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, code int, err error) {
	e := RespError{
		Err:     err,
		Message: err.Error(),
	}
	c.JSON(code, gin.H{
		"code":    code,
		"message": e.Message,
		"data":    e,
	})
}

func ErrorBadRequestResponse(c *gin.Context, err error) {
	e := RespError{
		Err:     err,
		Message: err.Error(),
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": e.Message,
		"data":    e,
	})
}

func ErrorInternalServerResponse(c *gin.Context, err error) {
	e := RespError{
		Err:     err,
		Message: err.Error(),
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": e.Message,
		"data":    e,
	})
}

func ErrorAuthenticationResponse(c *gin.Context, err error) {
	e := RespError{
		Err:     err,
		Message: err.Error(),
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": e.Message,
		"data":    e,
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

func SuccessResponseWithMsg(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}

func ApiResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
