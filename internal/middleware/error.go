package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	apperrors "github.com/ultimatum/apihub_go/pkg/errors"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// ErrorHandler returns an error handling middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Check if it's an AppError
			var appErr *apperrors.AppError
			if errors.As(err, &appErr) {
				response.Error(c, appErr.StatusCode, appErr.Message)
				return
			}

			// Default to internal server error
			response.InternalServerError(c, "An unexpected error occurred")
		}
	}
}
