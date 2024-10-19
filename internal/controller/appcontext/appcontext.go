package appcontext

import (
	"CodeSolveLearn_API/db"
	"CodeSolveLearn_API/internal/app_error"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppContext struct {
	*gin.Context
	DB db.Database
}

func NewHandler(handler func(ctx *AppContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := &AppContext{Context: c}
		handler(appCtx)
	}
}

func (c *AppContext) HandleResponse(obj interface{}, err error) {
	// check error type
	var httpError app_error.HttpError

	switch {
	case errors.As(err, &httpError):
		if httpError.Err == nil {
			statusCode := httpError.StatusCode

			// zero value statusCode implies a 200
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			c.JSON(statusCode, obj)
			return
		}
		_ = c.AbortWithError(httpError.StatusCode, httpError).SetType(gin.ErrorTypePrivate)
		return
	}

	// check error value
	switch {
	case err == nil:
		c.JSON(http.StatusOK, obj)
	case errors.Is(err, sql.ErrNoRows):
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}
